package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt" // Modul Enkripsi Password
)

var DB *mongo.Database

// --- MODEL DATA ---

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name       string             `bson:"name" json:"name"`
	Email      string             `bson:"email" json:"email"`
	Password   string             `bson:"password" json:"-"`                        // Sembunyikan password di respons JSON
	ProfilePic string             `bson:"profile_pic,omitempty" json:"profile_pic"` // Tambahkan baris ini
	Deposit    int                `bson:"deposit" json:"deposit"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
}

type Event struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title           string             `bson:"title" json:"title"`
	Location        string             `bson:"location" json:"location"`
	MatchDate       time.Time          `bson:"match_date" json:"match_date"`
	QuotaMax        int                `bson:"quota_max" json:"quota_max"`
	PricePerPerson  int                `bson:"price_per_person" json:"price_per_person"`
	PaymentDeadline time.Time          `bson:"payment_deadline" json:"payment_deadline"`
	Status          string             `bson:"status" json:"status"`
	CreatedAt       time.Time          `bson:"created_at" json:"created_at"`
}

type Registration struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	EventID         primitive.ObjectID `bson:"event_id" json:"event_id"`
	UserID          primitive.ObjectID `bson:"user_id" json:"user_id"` // Relasi ke tabel User
	UserName        string             `bson:"user_name" json:"user_name"`
	PollingStatus   string             `bson:"polling_status" json:"polling_status"`
	PaymentStatus   string             `bson:"payment_status" json:"payment_status"`
	PaymentMethod   string             `bson:"payment_method,omitempty" json:"payment_method"`
	PaymentProofURL string             `bson:"payment_proof_url,omitempty" json:"payment_proof_url"`
	RegisteredAt    time.Time          `bson:"registered_at" json:"registered_at"`
}

type Notification struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title     string             `bson:"title" json:"title"`
	Message   string             `bson:"message" json:"message"`
	Type      string             `bson:"type" json:"type"` // INFO, JADWAL, URGENT, PROMO
	Sender    string             `bson:"sender" json:"sender"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

// --- FUNGSI KONEKSI ---

func ConnectDB() {
	// Coba load file .env jika tersedia (tidak wajib jika variabel di-pass dari Docker)
	_ = godotenv.Load()

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "fun_football_db"
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Gagal koneksi ke MongoDB: ", err)
	}

	DB = client.Database(dbName)
	fmt.Println("🚀 Berhasil terhubung ke database:", dbName, "pada", mongoURI)
}

// --- HANDLER API AUTENTIKASI ---

func RegisterUser(c *gin.Context) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Cek apakah email sudah terdaftar sebelumnya
	count, _ := DB.Collection("users").CountDocuments(ctx, bson.M{"email": input.Email})
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email sudah terdaftar!"})
		return
	}

	// 1. Enkripsi Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal enkripsi password"})
		return
	}

	// 2. Siapkan data User
	newUser := User{
		Name:       input.Name,
		Email:      input.Email,
		Password:   string(hashedPassword),
		ProfilePic: "", // Kosongkan dulu saat pertama daftar
	}

	// 3. Simpan ke database
	res, err := DB.Collection("users").InsertOne(ctx, newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendaftar user"})
		return
	}
	newUser.ID = res.InsertedID.(primitive.ObjectID)

	c.JSON(http.StatusCreated, gin.H{"message": "Akun berhasil dibuat", "data": newUser})
}

func LoginUser(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 1. Cari user berdasarkan email
	var user User
	err := DB.Collection("users").FindOne(ctx, bson.M{"email": input.Email}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email tidak terdaftar"})
		return
	}

	// 2. Cocokkan password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password salah"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login berhasil", "data": user})
}

// --- HANDLER API LAINNYA ---

func CreateEvent(c *gin.Context) {
	var newEvent Event
	if err := c.ShouldBindJSON(&newEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newEvent.Status = "OPEN"
	newEvent.CreatedAt = time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := DB.Collection("events").InsertOne(ctx, newEvent)
	if err != nil {
		fmt.Println("❌ ERROR MONGODB:", err)                                // <-- Tambahkan baris ini untuk mencetak error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // <-- Tampilkan error asli ke frontend sementara
		return
	}
	newEvent.ID = result.InsertedID.(primitive.ObjectID)
	BroadcastEvent("EVENT_UPDATED", newEvent)
	c.JSON(http.StatusCreated, gin.H{"message": "Jadwal berhasil dibuat!", "data": newEvent})
}

type EventDetail struct {
	ID              primitive.ObjectID `json:"id"`
	Title           string             `json:"title"`
	Location        string             `json:"location"`
	MatchDate       time.Time          `json:"match_date"`
	QuotaMax        int                `json:"quota_max"`
	PricePerPerson  int                `json:"price_per_person"`
	PaymentDeadline time.Time          `json:"payment_deadline"`
	Status          string             `json:"status"`
	CreatedAt       time.Time          `json:"created_at"`
	RegisteredCount int                `json:"registered_count"`
	PaidCount       int                `json:"paid_count"`
	SlotsLeft       int                `json:"slots_left"`
}

func GetEvents(c *gin.Context) {
	var results []EventDetail
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Find().SetSort(bson.D{{Key: "match_date", Value: 1}})
	cursor, err := DB.Collection("events").Find(ctx, bson.M{}, opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var event Event
		cursor.Decode(&event)

		regCount, _ := DB.Collection("registrations").CountDocuments(ctx, bson.M{
			"event_id":       event.ID,
			"polling_status": "JOIN",
		})

		paidCount, _ := DB.Collection("registrations").CountDocuments(ctx, bson.M{
			"event_id":       event.ID,
			"payment_status": "PAID",
		})

		slotsLeft := event.QuotaMax - int(regCount)
		if slotsLeft < 0 {
			slotsLeft = 0
		}

		results = append(results, EventDetail{
			ID:              event.ID,
			Title:           event.Title,
			Location:        event.Location,
			MatchDate:       event.MatchDate,
			QuotaMax:        event.QuotaMax,
			PricePerPerson:  event.PricePerPerson,
			PaymentDeadline: event.PaymentDeadline,
			Status:          event.Status,
			CreatedAt:       event.CreatedAt,
			RegisteredCount: int(regCount),
			PaidCount:       int(paidCount),
			SlotsLeft:       slotsLeft,
		})
	}
	c.JSON(http.StatusOK, gin.H{"data": results})
}

func RegisterEvent(c *gin.Context) {
	var req struct {
		EventID string `json:"event_id"`
		UserID  string `json:"user_id"` // Berubah: Meminta UserID, bukan UserName lagi
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data salah"})
		return
	}

	eventObjID, _ := primitive.ObjectIDFromHex(req.EventID)
	userObjID, _ := primitive.ObjectIDFromHex(req.UserID)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Cek Data User untuk mengambil Nama
	var user User
	err := DB.Collection("users").FindOne(ctx, bson.M{"_id": userObjID}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Akun tidak ditemukan"})
		return
	}

	var event Event
	err = DB.Collection("events").FindOne(ctx, bson.M{"_id": eventObjID}).Decode(&event)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Jadwal tidak ditemukan"})
		return
	}

	var existingReg Registration
	errCheck := DB.Collection("registrations").FindOne(ctx, bson.M{
		"event_id": eventObjID,
		"user_id":  userObjID,
	}).Decode(&existingReg)

	if errCheck == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Anda sudah terdaftar di jadwal ini!"})
		return
	}

	count, _ := DB.Collection("registrations").CountDocuments(ctx, bson.M{
		"event_id":       eventObjID,
		"polling_status": "JOIN",
	})

	pollingStatus := "JOIN"
	if int(count) >= event.QuotaMax {
		pollingStatus = "WAITING_LIST"
	}

	newReg := Registration{
		EventID:       eventObjID,
		UserID:        userObjID,
		UserName:      user.Name, // Nama otomatis diambil dari database akun
		PollingStatus: pollingStatus,
		PaymentStatus: "UNPAID",
		RegisteredAt:  time.Now(),
	}

	res, err := DB.Collection("registrations").InsertOne(ctx, newReg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendaftar"})
		return
	}
	newReg.ID = res.InsertedID.(primitive.ObjectID)
	BroadcastEvent("REGISTRATION_UPDATED", newReg)
	c.JSON(http.StatusCreated, gin.H{"message": "Pendaftaran berhasil dicatat", "status_kuota": pollingStatus, "data": newReg})
}

func UploadPaymentProof(c *gin.Context) {
	regID := c.PostForm("registration_id")
	if regID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "registration_id tidak boleh kosong"})
		return
	}
	file, err := c.FormFile("receipt")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gagal mengambil file gambar"})
		return
	}
	filename := fmt.Sprintf("%d-%s", time.Now().Unix(), filepath.Base(file.Filename))
	uploadPath := fmt.Sprintf("uploads/%s", filename)

	if err := c.SaveUploadedFile(file, uploadPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan gambar"})
		return
	}
	regObjID, _ := primitive.ObjectIDFromHex(regID)
	update := bson.M{
		"$set": bson.M{
			"payment_status":    "VERIFYING",
			"payment_proof_url": "/" + uploadPath,
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = DB.Collection("registrations").UpdateOne(ctx, bson.M{"_id": regObjID}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal update database"})
		return
	}
	BroadcastEvent("PAYMENT_UPDATED", gin.H{"registration_id": regID, "url": "/" + uploadPath})
	c.JSON(http.StatusOK, gin.H{"message": "Bukti berhasil diunggah", "url": "/" + uploadPath})
}

func GetRegistrations(c *gin.Context) {
	var registrations []Registration
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{}
	eventIDStr := c.Query("event_id")
	if eventIDStr != "" {
		if objID, err := primitive.ObjectIDFromHex(eventIDStr); err == nil {
			filter["event_id"] = objID
		}
	}

	cursor, err := DB.Collection("registrations").Find(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var reg Registration
		cursor.Decode(&reg)
		registrations = append(registrations, reg)
	}
	c.JSON(http.StatusOK, gin.H{"data": registrations})
}

func VerifyPayment(c *gin.Context) {
	var req struct {
		RegistrationID string `json:"registration_id"`
		Action         string `json:"action"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data salah"})
		return
	}
	regObjID, err := primitive.ObjectIDFromHex(req.RegistrationID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID Pendaftaran tidak valid"})
		return
	}

	var newPaymentStatus string
	if req.Action == "APPROVE" {
		newPaymentStatus = "PAID"
	} else if req.Action == "REJECT" {
		newPaymentStatus = "REJECTED"
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Action harus APPROVE atau REJECT"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	update := bson.M{"$set": bson.M{"payment_status": newPaymentStatus}}
	_, err = DB.Collection("registrations").UpdateOne(ctx, bson.M{"_id": regObjID}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal update status"})
		return
	}
	BroadcastEvent("PAYMENT_UPDATED", gin.H{"registration_id": req.RegistrationID, "status": newPaymentStatus})
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Berhasil di-%s", req.Action), "status_baru": newPaymentStatus})
}

func GetMyRegistrations(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id diperlukan"})
		return
	}
	userObjID, _ := primitive.ObjectIDFromHex(userID)

	var registrations []Registration
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := DB.Collection("registrations").Find(ctx, bson.M{"user_id": userObjID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var reg Registration
		cursor.Decode(&reg)
		registrations = append(registrations, reg)
	}
	c.JSON(http.StatusOK, gin.H{"data": registrations})
}

func GetUser(c *gin.Context) {
	id := c.Query("id")
	objID, _ := primitive.ObjectIDFromHex(id)

	var user User
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := DB.Collection("users").FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateProfile(c *gin.Context) {
	_ = c.Request.ParseMultipartForm(32 << 20)
	userID := c.PostForm("user_id")
	newName := c.PostForm("name")

	if userID == "" || newName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak lengkap"})
		return
	}

	objID, _ := primitive.ObjectIDFromHex(userID)
	updateData := bson.M{"name": newName}

	var photoURL string
	// Cek apakah ada file foto yang diunggah
	file, err := c.FormFile("profile_pic")
	if err == nil && file != nil {
		os.MkdirAll("uploads", os.ModePerm)
		ext := filepath.Ext(file.Filename)
		if ext == "" {
			ext = ".jpg"
		}
		filename := fmt.Sprintf("avatar-%d%s", time.Now().UnixNano(), ext)
		uploadPath := fmt.Sprintf("uploads/%s", filename)
		if err := c.SaveUploadedFile(file, uploadPath); err != nil {
			log.Println("Gagal menyimpan file avatar:", err)
		} else {
			photoURL = "/" + uploadPath
			updateData["profile_pic"] = photoURL
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = DB.Collection("users").UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": updateData})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal update profile"})
		return
	}

	BroadcastEvent("USER_UPDATED", gin.H{"user_id": userID, "new_name": newName, "profile_pic": photoURL})
	c.JSON(http.StatusOK, gin.H{"message": "Profile berhasil diperbarui", "new_name": newName, "profile_pic": photoURL})
}

func GetFinancialSummary(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 1. Ambil semua pendaftaran yang status pembayarannya PAID (Lunas)
	cursor, err := DB.Collection("registrations").Find(ctx, bson.M{"payment_status": "PAID"})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data keuangan"})
		return
	}
	defer cursor.Close(ctx)

	var totalKas int = 0
	var userContributions []gin.H

	userIDQuery := c.Query("user_id")
	var userObjID primitive.ObjectID
	if userIDQuery != "" {
		userObjID, _ = primitive.ObjectIDFromHex(userIDQuery)
	}

	for cursor.Next(ctx) {
		var reg Registration
		cursor.Decode(&reg)

		// Cari informasi event terkait untuk mendapatkan harga patungan
		var event Event
		errEvt := DB.Collection("events").FindOne(ctx, bson.M{"_id": reg.EventID}).Decode(&event)

		if errEvt == nil {
			totalKas += event.PricePerPerson // Akumulasi total kas komunitas

			// Jika ini adalah data milik user yang sedang login, masukkan ke riwayat kontribusinya
			if userIDQuery != "" && reg.UserID == userObjID {
				userContributions = append(userContributions, gin.H{
					"event_title": event.Title,
					"amount":      event.PricePerPerson,
					"date":        reg.RegisteredAt,
				})
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"total_kas":        totalKas,
		"my_contributions": userContributions,
	})
}

func GetDashboardStats(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 1. Hitung total jadwal
	totalEvents, _ := DB.Collection("events").CountDocuments(ctx, bson.M{})

	// 2. Hitung total pemain yang status polling-nya 'JOIN'
	totalPlayers, _ := DB.Collection("registrations").CountDocuments(ctx, bson.M{"polling_status": "JOIN"})

	// 3. Hitung total yang menunggu verifikasi (status pembayaran 'VERIFYING')
	pendingVerification, _ := DB.Collection("registrations").CountDocuments(ctx, bson.M{"payment_status": "VERIFYING"})

	c.JSON(http.StatusOK, gin.H{
		"total_events":         totalEvents,
		"total_players":        totalPlayers,
		"pending_verification": pendingVerification,
	})
}

func UpdateEvent(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var input Event
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data salah"})
		return
	}

	updateData := bson.M{
		"title":            input.Title,
		"location":         input.Location,
		"match_date":       input.MatchDate,
		"quota_max":        input.QuotaMax,
		"price_per_person": input.PricePerPerson,
		"payment_deadline": input.PaymentDeadline,
		"status":           input.Status,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = DB.Collection("events").UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": updateData})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui jadwal"})
		return
	}

	BroadcastEvent("EVENT_UPDATED", gin.H{"id": id, "action": "update"})
	c.JSON(http.StatusOK, gin.H{"message": "Jadwal berhasil diperbarui"})
}

func DeleteEvent(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = DB.Collection("events").DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus jadwal"})
		return
	}

	BroadcastEvent("EVENT_UPDATED", gin.H{"id": id, "action": "delete"})
	c.JSON(http.StatusOK, gin.H{"message": "Jadwal berhasil dihapus"})
}

// Tambahkan fungsi ini di atas atau di bawah fungsi API lainnya
func RequestTopUp(c *gin.Context) {
	// 1. Ambil data teks dari form-data
	userIdStr := c.PostForm("user_id")
	amountStr := c.PostForm("amount")

	if userIdStr == "" || amountStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID dan Nominal harus diisi"})
		return
	}

	// Ubah User ID menjadi format ObjectID MongoDB
	objID, err := primitive.ObjectIDFromHex(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID tidak valid"})
		return
	}

	// Ubah Amount (string) menjadi angka (integer)
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nominal top up tidak valid"})
		return
	}

	// 2. Tangani file gambar (Bukti Transfer)
	file, err := c.FormFile("receipt")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bukti transfer wajib diunggah"})
		return
	}

	// Buat folder penyimpanan khusus top up jika belum ada
	os.MkdirAll("./uploads/topups", os.ModePerm)

	// Buat nama file unik menggunakan waktu (timestamp) agar tidak bentrok
	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
	filepath := "./uploads/topups/" + filename
	fileUrl := "/uploads/topups/" + filename

	// Simpan file ke dalam folder komputer/server
	if err := c.SaveUploadedFile(file, filepath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan foto bukti transfer"})
		return
	}

	// 3. Siapkan data untuk disimpan ke koleksi "topups" di MongoDB
	newTopUp := bson.M{
		"user_id":    objID,
		"amount":     amount,
		"receipt":    fileUrl,
		"status":     "PENDING", // Status awal selalu PENDING, menunggu admin klik "Setujui"
		"created_at": time.Now(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Simpan ke database
	_, err = DB.Collection("topups").InsertOne(ctx, newTopUp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan permintaan deposit ke database"})
		return
	}

	BroadcastEvent("TOPUP_UPDATED", gin.H{"user_id": userIdStr, "amount": amount})
	c.JSON(http.StatusOK, gin.H{"message": "Permintaan deposit berhasil dikirim"})
}

func ApproveTopUp(c *gin.Context) {
	var input struct {
		TopUpID string `json:"topup_id"`
		Action  string `json:"action"` // "APPROVE" atau "REJECT"
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid"})
		return
	}

	topupObjID, err := primitive.ObjectIDFromHex(input.TopUpID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID TopUp tidak valid"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 1. Ambil data topup berdasarkan ID
	var topup struct {
		ID     primitive.ObjectID `bson:"_id"`
		UserID primitive.ObjectID `bson:"user_id"`
		Amount int                `bson:"amount"`
		Status string             `bson:"status"`
	}

	err = DB.Collection("topups").FindOne(ctx, bson.M{"_id": topupObjID}).Decode(&topup)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Permintaan top up tidak ditemukan"})
		return
	}

	if topup.Status != "PENDING" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Permintaan ini sudah diproses sebelumnya"})
		return
	}

	// 2. Jika disetujui (APPROVE), tambahkan saldo ke user
	if input.Action == "APPROVE" {
		// Update status topup jadi APPROVED
		_, err = DB.Collection("topups").UpdateOne(ctx, bson.M{"_id": topupObjID}, bson.M{"$set": bson.M{"status": "APPROVED"}})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui status topup"})
			return
		}

		// Tambahkan nominal deposit ke akun user ($inc)
		_, err = DB.Collection("users").UpdateOne(ctx, bson.M{"_id": topup.UserID}, bson.M{"$inc": bson.M{"deposit": topup.Amount}})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambahkan saldo ke user"})
			return
		}

		BroadcastEvent("TOPUP_UPDATED", gin.H{"topup_id": input.TopUpID, "action": "APPROVE"})
		c.JSON(http.StatusOK, gin.H{"message": "Top up berhasil disetujui dan saldo ditambahkan"})
		return
	}

	// 3. Jika ditolak (REJECT)
	if input.Action == "REJECT" {
		_, err = DB.Collection("topups").UpdateOne(ctx, bson.M{"_id": topupObjID}, bson.M{"$set": bson.M{"status": "REJECTED"}})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menolak topup"})
			return
		}

		BroadcastEvent("TOPUP_UPDATED", gin.H{"topup_id": input.TopUpID, "action": "REJECT"})
		c.JSON(http.StatusOK, gin.H{"message": "Permintaan top up ditolak"})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "Aksi tidak dikenali"})
}

func GetPendingTopUps(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Menggabungkan (join) koleksi topups dengan users untuk mendapatkan nama pemain
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.M{"status": "PENDING"}}},
		{{Key: "$lookup", Value: bson.M{
			"from":         "users",
			"localField":   "user_id",
			"foreignField": "_id",
			"as":           "user_info",
		}}},
		{{Key: "$unwind", Value: bson.M{"path": "$user_info", "preserveNullAndEmptyArrays": true}}},
		{{Key: "$sort", Value: bson.M{"created_at": -1}}}, // Urutkan dari yang terbaru
		{{Key: "$project", Value: bson.M{
			"_id":        1,
			"amount":     1,
			"receipt":    1,
			"status":     1,
			"created_at": 1,
			"user_name":  "$user_info.name",
		}}},
	}

	cursor, err := DB.Collection("topups").Aggregate(ctx, pipeline)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data top up"})
		return
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membaca data"})
		return
	}

	// Jika kosong, kembalikan array kosong agar frontend tidak error
	if results == nil {
		results = []bson.M{}
	}

	c.JSON(http.StatusOK, gin.H{"data": results})
}

func GetApprovedTopUps(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Ambil data top up yang sudah APPROVED dan gabungkan dengan nama user
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.M{"status": "APPROVED"}}},
		{{Key: "$lookup", Value: bson.M{
			"from":         "users",
			"localField":   "user_id",
			"foreignField": "_id",
			"as":           "user_info",
		}}},
		{{Key: "$unwind", Value: bson.M{"path": "$user_info", "preserveNullAndEmptyArrays": true}}},
		{{Key: "$sort", Value: bson.M{"created_at": -1}}},
		{{Key: "$project", Value: bson.M{
			"_id":        1,
			"amount":     1,
			"status":     1,
			"created_at": 1,
			"user_name":  "$user_info.name",
		}}},
	}

	cursor, err := DB.Collection("topups").Aggregate(ctx, pipeline)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data kas deposit"})
		return
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membaca data kas deposit"})
		return
	}

	if results == nil {
		results = []bson.M{}
	}

	c.JSON(http.StatusOK, gin.H{"data": results})
}

func PayDeposit(c *gin.Context) {
	var input struct {
		RegistrationID string `json:"registration_id"`
		UserID         string `json:"user_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data input tidak valid"})
		return
	}

	regObjID, err := primitive.ObjectIDFromHex(input.RegistrationID)
	userObjID, err2 := primitive.ObjectIDFromHex(input.UserID)
	if err != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 1. Cari data pendaftaran (tiket)
	var reg struct {
		ID      primitive.ObjectID `bson:"_id"`
		EventID primitive.ObjectID `bson:"event_id"`
		Status  string             `bson:"payment_status"`
	}
	err = DB.Collection("registrations").FindOne(ctx, bson.M{"_id": regObjID}).Decode(&reg)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data pendaftaran tidak ditemukan"})
		return
	}

	if reg.Status == "PAID" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Pendaftaran ini sudah lunas"})
		return
	}

	// 2. Cari data jadwal (untuk mengetahui harganya)
	var event struct {
		ID    primitive.ObjectID `bson:"_id"`
		Price int                `bson:"price_per_person"`
	}
	err = DB.Collection("events").FindOne(ctx, bson.M{"_id": reg.EventID}).Decode(&event)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data jadwal tidak ditemukan"})
		return
	}

	// 3. Cari data pemain (untuk cek saldo saat ini)
	var user struct {
		ID      primitive.ObjectID `bson:"_id"`
		Deposit int                `bson:"deposit"`
	}
	err = DB.Collection("users").FindOne(ctx, bson.M{"_id": userObjID}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data user tidak ditemukan"})
		return
	}

	// 4. Validasi: Apakah saldo cukup?
	if user.Deposit < event.Price {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Saldo deposit tidak mencukupi!"})
		return
	}

	// 5. Potong saldo pemain
	_, err = DB.Collection("users").UpdateOne(
		ctx,
		bson.M{"_id": userObjID},
		bson.M{"$inc": bson.M{"deposit": -event.Price}}, // Gunakan minus untuk memotong
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memotong saldo deposit"})
		return
	}

	// 6. Ubah status tiket langsung menjadi PAID (Tanpa perlu verifikasi admin)
	_, err = DB.Collection("registrations").UpdateOne(
		ctx,
		bson.M{"_id": regObjID},
		bson.M{"$set": bson.M{
			"payment_status": "PAID",
			"payment_method": "deposit",
		}},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengupdate status tiket"})
		return
	}

	BroadcastEvent("PAYMENT_UPDATED", gin.H{"registration_id": input.RegistrationID, "user_id": input.UserID})
	c.JSON(http.StatusOK, gin.H{"message": "Berhasil bayar pakai deposit! 🔥"})
}

// 1. Mengambil Profil Admin
func GetAdminProfile(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var admin bson.M
	// Mencari user yang memiliki field role: "admin"
	err := DB.Collection("users").FindOne(ctx, bson.M{"role": "admin"}).Decode(&admin)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data admin tidak ditemukan di database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    admin["_id"],
		"name":  admin["name"],
		"email": admin["email"],
	})
}

// 2. Memperbarui Nama & Email
func UpdateAdminProfile(c *gin.Context) {
	var input struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid"})
		return
	}

	objID, _ := primitive.ObjectIDFromHex(input.ID)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := DB.Collection("users").UpdateOne(
		ctx,
		bson.M{"_id": objID},
		bson.M{"$set": bson.M{"name": input.Name, "email": input.Email}},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan profil"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Profil berhasil diperbarui"})
}

// 3. Memperbarui Password
func UpdateAdminPassword(c *gin.Context) {
	var input struct {
		ID       string `json:"id"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid"})
		return
	}

	objID, _ := primitive.ObjectIDFromHex(input.ID)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Catatan: Di production, password sangat disarankan untuk di-hash (misal dengan bcrypt)
	_, err := DB.Collection("users").UpdateOne(
		ctx,
		bson.M{"_id": objID},
		bson.M{"$set": bson.M{"password": input.Password}},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengupdate password"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Password berhasil diperbarui"})
}

// --- REALTIME WEBSOCKET HUB ---

type WSMessage struct {
	Type    string      `json:"type"`    // Contoh: "REGISTRATION_UPDATED", "PAYMENT_UPDATED", "EVENT_UPDATED", "TOPUP_UPDATED"
	Payload interface{} `json:"payload"` // Data detail objek
}

type WSHub struct {
	clients    map[*websocket.Conn]bool
	broadcast  chan WSMessage
	register   chan *websocket.Conn
	unregister chan *websocket.Conn
	mutex      sync.RWMutex
}

var Hub = &WSHub{
	clients:    make(map[*websocket.Conn]bool),
	broadcast:  make(chan WSMessage, 100),
	register:   make(chan *websocket.Conn),
	unregister: make(chan *websocket.Conn),
}

func (h *WSHub) Run() {
	for {
		select {
		case conn := <-h.register:
			h.mutex.Lock()
			h.clients[conn] = true
			h.mutex.Unlock()
			fmt.Println("🔌 Client WebSocket Terhubung. Total aktif:", len(h.clients))

		case conn := <-h.unregister:
			h.mutex.Lock()
			if _, ok := h.clients[conn]; ok {
				delete(h.clients, conn)
				conn.Close()
			}
			h.mutex.Unlock()
			fmt.Println("🔌 Client WebSocket Terputus. Total aktif:", len(h.clients))

		case msg := <-h.broadcast:
			h.mutex.RLock()
			for conn := range h.clients {
				err := conn.WriteJSON(msg)
				if err != nil {
					conn.Close()
					delete(h.clients, conn)
				}
			}
			h.mutex.RUnlock()
		}
	}
}

// Fungsi pembantu untuk mengirim notifikasi realtime ke seluruh client
func BroadcastEvent(eventType string, payload interface{}) {
	Hub.broadcast <- WSMessage{
		Type:    eventType,
		Payload: payload,
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Izinkan semua origin terhubung ke WebSocket
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Gagal upgrade WebSocket:", err)
		return
	}

	Hub.register <- conn

	go func() {
		defer func() {
			Hub.unregister <- conn
		}()
		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				break
			}
		}
	}()
}

// --- MIDDLEWARE CORS ---
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		}
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, ngrok-skip-browser-warning, *")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE, HEAD")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// --- HANDLER NOTIFIKASI & BROADCAST ---

func BroadcastNotification(c *gin.Context) {
	var input struct {
		Title   string `json:"title"`
		Message string `json:"message"`
		Type    string `json:"type"`
		Sender  string `json:"sender"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid"})
		return
	}

	if input.Title == "" || input.Message == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Judul dan isi pesan wajib diisi"})
		return
	}

	if input.Type == "" {
		input.Type = "INFO"
	}
	if input.Sender == "" {
		input.Sender = "Admin JMT Sport"
	}

	notif := Notification{
		ID:        primitive.NewObjectID(),
		Title:     input.Title,
		Message:   input.Message,
		Type:      input.Type,
		Sender:    input.Sender,
		CreatedAt: time.Now(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := DB.Collection("notifications").InsertOne(ctx, notif)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data notifikasi"})
		return
	}

	// Broadcast seketika ke seluruh pemain via WebSocket
	BroadcastEvent("BROADCAST_NOTIFICATION", notif)

	c.JSON(http.StatusOK, gin.H{
		"message": "Notifikasi berhasil disiarkan ke seluruh aplikasi!",
		"data":    notif,
	})
}

func GetNotifications(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}}).SetLimit(50)
	cursor, err := DB.Collection("notifications").Find(ctx, bson.M{}, opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil daftar notifikasi"})
		return
	}
	defer cursor.Close(ctx)

	var list []Notification
	if err := cursor.All(ctx, &list); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memproses data notifikasi"})
		return
	}

	if list == nil {
		list = []Notification{}
	}

	c.JSON(http.StatusOK, list)
}

func DeleteNotification(c *gin.Context) {
	idParam := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID notifikasi tidak valid"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = DB.Collection("notifications").DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus notifikasi"})
		return
	}

	BroadcastEvent("NOTIFICATION_DELETED", gin.H{"id": idParam})
	c.JSON(http.StatusOK, gin.H{"message": "Notifikasi berhasil dihapus"})
}

// --- MAIN FUNGSI ---

func main() {
	ConnectDB()
	os.MkdirAll("uploads", os.ModePerm)

	// Jalankan Realtime WebSocket Hub di background goroutine
	go Hub.Run()

	r := gin.Default()

	// Pasang Middleware CORS
	r.Use(CORSMiddleware())

	// Endpoint Realtime WebSocket
	r.GET("/ws", HandleWebSocket)

	r.Static("/uploads", "./uploads")

	r.POST("/request-topup", RequestTopUp)
	r.POST("/approve-topup", ApproveTopUp)
	r.GET("/pending-topups", GetPendingTopUps)
	// Endpoint Auth (Baru)
	r.POST("/auth/register", RegisterUser)
	r.POST("/auth/login", LoginUser)

	// Endpoint Events & Registrations
	r.POST("/events", CreateEvent)
	r.GET("/events", GetEvents)
	r.POST("/register", RegisterEvent)
	r.POST("/upload-proof", UploadPaymentProof)
	r.GET("/registrations", GetRegistrations)
	r.PUT("/verify-payment", VerifyPayment)
	r.GET("/approved-topups", GetApprovedTopUps)
	r.POST("/pay-deposit", PayDeposit)

	r.GET("/user", GetUser)
	r.PUT("/update-profile", UpdateProfile)
	r.POST("/update-profile", UpdateProfile)

	r.GET("/my-registrations", GetMyRegistrations) // Tambahkan baris ini

	r.GET("/financial-summary", GetFinancialSummary)
	r.GET("/dashboard-stats", GetDashboardStats)

	r.PUT("/events/:id", UpdateEvent)    // Endpoint Edit
	r.DELETE("/events/:id", DeleteEvent) // Endpoint Hapus
	r.POST("/register-user", RegisterUser)
	r.POST("/login", LoginUser)

	r.GET("/admin-profile", GetAdminProfile)
	r.PUT("/admin-profile", UpdateAdminProfile)
	r.PUT("/admin-password", UpdateAdminPassword)

	// Endpoint Notifikasi & Broadcast
	r.POST("/broadcast-notification", BroadcastNotification)
	r.GET("/notifications", GetNotifications)
	r.DELETE("/notifications/:id", DeleteNotification)

	// Frontend SPA & Nuxt Proxy Handler (Mencegah 404 pada route /player, /admin, dll)
	r.NoRoute(func(c *gin.Context) {
		reqPath := c.Request.URL.Path

		// Jangan handle rute WebSocket
		if reqPath == "/ws" {
			c.Status(http.StatusNotFound)
			return
		}

		publicDir := filepath.Join("fun-football-admin", ".output", "public")

		// 1. Jika ada file statis langsung di folder .output/public (misal: logo-jmt.png, manifest.webmanifest, _nuxt/*)
		targetFile := filepath.Join(publicDir, filepath.Clean(reqPath))
		if stat, err := os.Stat(targetFile); err == nil && !stat.IsDir() {
			c.File(targetFile)
			return
		}

		// 2. Cek apakah ada file HTML spesifik hasil prerender (misal: player.html atau player/index.html)
		htmlFile := filepath.Join(publicDir, filepath.Clean(reqPath)+".html")
		if stat, err := os.Stat(htmlFile); err == nil && !stat.IsDir() {
			c.File(htmlFile)
			return
		}
		indexHtmlFile := filepath.Join(publicDir, filepath.Clean(reqPath), "index.html")
		if stat, err := os.Stat(indexHtmlFile); err == nil && !stat.IsDir() {
			c.File(indexHtmlFile)
			return
		}

		// 3. Proxy ke Nuxt Dev Server (localhost:3000) jika sedang aktif
		targetURL, _ := url.Parse("http://localhost:3000")
		proxy := httputil.NewSingleHostReverseProxy(targetURL)
		proxy.ErrorHandler = func(w http.ResponseWriter, req *http.Request, err error) {
			// Jika Nuxt Dev server tidak aktif, fallback ke index.html / 200.html bawaan SPA
			spa200 := filepath.Join(publicDir, "200.html")
			if _, err200 := os.Stat(spa200); err200 == nil {
				http.ServeFile(w, req, spa200)
				return
			}
			spaIndex := filepath.Join(publicDir, "index.html")
			if _, errIndex := os.Stat(spaIndex); errIndex == nil {
				http.ServeFile(w, req, spaIndex)
				return
			}
			http.Error(w, "Halaman tidak ditemukan", http.StatusNotFound)
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Server berjalan di port :", port)
	r.Run(":" + port)
}

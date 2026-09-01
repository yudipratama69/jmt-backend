package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"path/filepath"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"-"` // Sembunyikan password di respons JSON
	ProfilePic string             `bson:"profile_pic,omitempty" json:"profile_pic"` // Tambahkan baris ini
	Deposit    int                `bson:"deposit" json:"deposit"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`

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

// --- FUNGSI KONEKSI ---

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DB_NAME")

	clientOptions := options.Client().ApplyURI(mongoURI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	DB = client.Database(dbName)
	fmt.Println("🚀 Berhasil terhubung ke database:", dbName)
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data"})
		return
	}
	newEvent.ID = result.InsertedID.(primitive.ObjectID)
	c.JSON(http.StatusCreated, gin.H{"message": "Jadwal berhasil dibuat!", "data": newEvent})
}

func GetEvents(c *gin.Context) {
	var events []Event
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := DB.Collection("events").Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var event Event
		cursor.Decode(&event)
		events = append(events, event)
	}
	c.JSON(http.StatusOK, gin.H{"data": events})
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
	c.JSON(http.StatusOK, gin.H{"message": "Bukti berhasil diunggah", "url": "/" + uploadPath})
}

func GetRegistrations(c *gin.Context) {
	var registrations []Registration
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := DB.Collection("registrations").Find(ctx, bson.M{})
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
	userID := c.PostForm("user_id")
	newName := c.PostForm("name")
	
	if userID == "" || newName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak lengkap"})
		return
	}

	objID, _ := primitive.ObjectIDFromHex(userID)
	updateData := bson.M{"name": newName}

	// Cek apakah ada file foto yang diunggah
	file, err := c.FormFile("profile_pic")
	if err == nil { // Jika ada file foto
		filename := fmt.Sprintf("avatar-%d-%s", time.Now().Unix(), filepath.Base(file.Filename))
		uploadPath := fmt.Sprintf("uploads/%s", filename)
		c.SaveUploadedFile(file, uploadPath)
		updateData["profile_pic"] = "/" + uploadPath
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = DB.Collection("users").UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": updateData})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal update profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile berhasil diperbarui", "new_name": newName})
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
		"total_kas":          totalKas,
		"my_contributions":   userContributions,
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
		"title":             input.Title,
		"location":          input.Location,
		"match_date":        input.MatchDate,
		"quota_max":         input.QuotaMax,
		"price_per_person":  input.PricePerPerson,
		"payment_deadline":  input.PaymentDeadline,
		"status":            input.Status,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = DB.Collection("events").UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": updateData})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui jadwal"})
		return
	}

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
		{{"$match", bson.M{"status": "PENDING"}}},
		{{"$lookup", bson.M{
			"from":         "users",
			"localField":   "user_id",
			"foreignField": "_id",
			"as":           "user_info",
		}}},
		{{"$unwind", bson.M{"path": "$user_info", "preserveNullAndEmptyArrays": true}}},
		{{"$sort", bson.M{"created_at": -1}}}, // Urutkan dari yang terbaru
		{{"$project", bson.M{
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
		{{"$match", bson.M{"status": "APPROVED"}}},
		{{"$lookup", bson.M{
			"from":         "users",
			"localField":   "user_id",
			"foreignField": "_id",
			"as":           "user_info",
		}}},
		{{"$unwind", bson.M{"path": "$user_info", "preserveNullAndEmptyArrays": true}}},
		{{"$sort", bson.M{"created_at": -1}}},
		{{"$project", bson.M{
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
// --- MAIN FUNGSI ---

func main() {
	ConnectDB()
	os.MkdirAll("uploads", os.ModePerm)

	r := gin.Default()
	r.Use(cors.Default())
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

	r.GET("/my-registrations", GetMyRegistrations) // Tambahkan baris ini

	r.GET("/financial-summary", GetFinancialSummary)
	r.GET("/dashboard-stats", GetDashboardStats)

	r.PUT("/events/:id", UpdateEvent)   // Endpoint Edit
	r.DELETE("/events/:id", DeleteEvent) // Endpoint Hapus
	r.POST("/register-user", RegisterUser)
	r.POST("/login", LoginUser)

	r.GET("/admin-profile", GetAdminProfile)
	r.PUT("/admin-profile", UpdateAdminProfile)
	r.PUT("/admin-password", UpdateAdminPassword)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Server berjalan di port :", port)
	r.Run(":" + port)
}
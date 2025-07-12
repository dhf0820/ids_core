package main

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	//"net/http"
	//"os"

	"strings"
	"time"

	//auth "github.com/dhf0820/authorize"
	//jwt "github.com/dhf0820/jwToken"
	jwt "github.com/dhf0820/golangJWT"
	common "github.com/dhf0820/uc_common"

	//"github.com/google/uuid"
	nano_uuid "github.com/aidarkhanov/nanoid/v2"
	"github.com/davecgh/go-spew/spew"

	//"github.com/dgrijalva/jwt-go"

	vsLog "github.com/dhf0820/vslog"
	//"github.com/sirupsen/logrus"
	fhir "github.com/dhf0820/fhir4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//var jwtKey = []byte(os.Getenv("ACCESS_SECRET"))
//var jwtRefreshKey = []byte(os.Getenv("REFRESH_SECRET"))

type User struct {
	ID                   primitive.ObjectID   `json:"id" bson:"_id"` //UserId Must be Email Address
	UserName             string               `json:"user_name" bson:"user_name"`
	Password             string               `json:"password" bson:"password"`
	FullName             string               `json:"full_name" bson:"full_name"`
	Phone                string               `json:"phone" bson:"phone"`
	Role                 string               `json:"role" bson:"role"` //Single String role: Provider, Admin, Nurse, Office
	LastLogin            *time.Time           `json:"last_login" bson:"last_login"`
	LastAttempt          *time.Time           `json:"last_attempt" bson:"last_attempt"`
	Attempts             int                  `json:"attempts" bson:"attempts"`
	PasswordChangedAt    *time.Time           `json:"passwordChangedAt" bson:"password_changed_at"`
	CreatedAt            *time.Time           `json:"created_at" bson:"created_at"`
	UpdatedAt            *time.Time           `json:"updated_at" bson:"updated_at"`
	DeletedAt            *time.Time           `json:"deleted_at" bson:"deleted_at"`
	CurrentLocalPatient  *ActivePatient       `json:"currentLocalPatient" bson:"currentLocalPatient"` //Current local patient
	CurrentRemotePatient *ActivePatient       `json:"currentRemotePatient" bson:"currentRemotePatient"`
	ResourceConfig       *fhir.ResourceConfig `json:"resourceConfig,omitempty"`
	Locals               []*Facility          `json:"locals" bson:"locals"`   //Medical Practice user is part of. Shows on Save
	Remotes              []*Facility          `json:"remotes" bson:"remotes"` //Medical Remotes user has access to, can include the emr for the Practice

}

type UserSummary struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	FullName string             `json:"fullName" bson:"fullName"`
}

// type User struct {
// 	ID          primitive.ObjectID `json:"id" bson:"_id"` //UserId Must be Email Address
// 	UserName    string             `json:"user_name" bson:"user_name"`
// 	Password    string             `json:"password" bson:"password"`
// 	FullName    string             `json:"full_name" bson:"full_name"`
// 	Phone       string             `json:"phone" bson:"phone"`
// 	Role        string             `json:"role" bson:"role"` //Single String role: Provider, Admin, Nurse, Office
// 	LastLogin   *time.Time         `json:"last_login" bson:"last_login"`
// 	LastAttempt *time.Time         `json:"last_attempt" bson:"last_attempt"`
// 	Attempts    int                `json:"attempts" bson:"attempts"`
// 	//Locals         []*Practice `json:"practices" bson:"practices"`   //Medical Practice user is part of. Shows on Save
// 	Remotes        []*Facility `json:"facilities" bson:"facilities"` //Medical Remotes user has access to, can include the emr for the Practice
// 	PasswordChangedAt *time.Time  `json:"passwordChangedAt" bson:"password_changed_at"`
// 	CreatedAt         *time.Time  `json:"created_at" bson:"created_at"`
// 	UpdatedAt         *time.Time  `json:"updated_at" bson:"updated_at"`
// 	DeletedAt         *time.Time  `json:"deleted_at" bson:"deleted_at"`
// }

// type Entity struct {
// 	Name       string      `json:"name" bson:"name"`
// 	IconLogo   string      `json:"iconLogo" bson:"iconLogo"`
// 	BaseUrl    string      `json:"baseUrl" bson:"baseUrl"`
// 	Remotes []*Facility `json:"facilities" bson:"facilities"`
// }

// func CreateAuth(userId string, td *TokenDetails) error {
// 	at := time.Unix(td.AtExpires, 0)
// 	rt := time.Unix(td.RtExpires, 0)

// }

func Login(userName, password, ip string) (*common.User, error) {
	var err error
	//loginTime := time.Now()
	// hpwd, err := HashPassword(password)
	// if err != nil {
	// 	return nil, nil, vsLog.Errorf("HashPassword failed: " + err.Error())
	// }
	//vsLog.Info("HashedPassword: " + hpwd)
	//pwd := EncryptPassword(password)
	//vsLog.Info("EncryptedPassword: " + pwd)

	//TODO: dd Practice to filter
	startTime := time.Now()
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return nil, vsLog.Errorf(fmt.Sprintf("UserName: %s Not Authorized  ERROR: %s", userName, err.Error()))
	}
	vsLog.Debug2(fmt.Sprintf("HashedPassword: %s", hashedPassword))
	//hashedPasswordByte, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	//filter := bson.D{bson.E{"user_name", userName}, {"password", hashedPassword}}
	filter := bson.M{"user_name": userName}
	//filter := bson.D{bson.E{Key: "user_name", Value: userName}, bson.E{Key: "password", Value: hashedPassword}}
	vsLog.Debug2(fmt.Sprintf("Calling FindOne:%s", filter))
	//filter := bson.D{{"user_name", userName}, {"password", pwd}}
	//filter := bson.D{bson.E{Key: "user_name", Value: userName}}
	//fmt.Printf("Using Filter: %v\n", filter)
	vsLog.Debug2("CallingGetCollection: user")
	collection, _ := GetCollection("user")
	usr := &common.User{}
	vsLog.Debug2(fmt.Sprintf("Calling FindOne: %v", filter))
	result := collection.FindOne(context.TODO(), filter)
	err = result.Err()
	if err != nil {
		vsLog.Error("FindOne error: " + err.Error())
		return nil, vsLog.Errorf(fmt.Sprintf("UserName: %s Not Authorized  ERROR: %s", userName, err.Error()))

	}
	vsLog.Debug2("No Error from FindOne")
	//err = collection.FindOne(context.TODO(), filter).Decode(&usr)
	// if result.Err != nil {
	// 	return nil, vsLog.Errorf(fmt.Sprintf("UserName: %s Not Authorized  ERROR: %s", userName, err.Error()))
	// }
	//vsLog.Debug2(fmt.Sprintf("result: %s", spew.Sdump(result)))
	err = result.Decode(&usr)
	if err != nil {
		vsLog.Error("Decode error: " + err.Error())
		return nil, vsLog.Errorf(fmt.Sprintf("UserName: %s Not Authorized  ERROR: %s", userName, err.Error()))
	}
	vsLog.Debug3(fmt.Sprintf("Login took %s", time.Since(startTime)))
	vsLog.Debug2(fmt.Sprintf("FindOneUser returned: %s", spew.Sdump(usr)))
	//fmt.Printf("Login:92 user- %s\n", user)
	startTime = time.Now()
	vsLog.Debug2("Calling CheckPassword")
	err = usr.CheckPassword(password)
	if err != nil {
		return nil, errors.New("not Authorized")
	}
	vsLog.Debug3(fmt.Sprintf("CheckPassword took %s", time.Since(startTime)))
	vsLog.Debug1(usr.UserName + " Logged in")
	return usr, nil
	/* 	tokenDuration, err := strconv.Atoi(os.Getenv("TOKEN_DURATION")) // In Minutes  default = 15
	   	if err != nil {
	   		vsLog.Warn("TOKEN_DURATION not set defaulting to: 60 min ")
	   		tokenDuration = 60
	   	}
	   	startTime = time.Now()
	   	duration := time.Duration(tokenDuration) * time.Minute
	   	jwt, err := CreateToken(ip, usr.UserName, duration, usr.ID.Hex(), usr.FullName, usr.Role, "")
	   	//token, err := CreateToken(ip, usr.ID.Hex(), usr.ClientId, duration, usr.FullName, usr.Role)
	   	vsLog.Info(fmt.Sprintf("CreateToken took: %s", time.Since(startTime)))
	   	//payload, ErrCode, err := ValidateToken(token, "") */
	// jwt, payload, err := jw_token.CreateToken(ip, user.UserName, duration, user.ID.Hex(), user.FullName, user.Role, as.ID.Hex())
	// if err != nil {
	// 	return vsLog.Errorf("Call to CreateJWToken failed: " + err.Error())
	// }
	// vsLog.Info("jwt: " + jwt)
	// vsLog.Info("payload: " + spew.Sdump(payload))
	//vsLog.Info("Calling auth.CreateSessionForUser")
	// AuthSession, err := auth.CreateSessionForUser(&usr, ip)
	// if err != nil {
	// 	return nil, nil, vsLog.Errorf("CreateSessionForUser failed: " + err.Error())
	// }
	// vsLog.Info(fmt.Sprintf("\nAuthSession: %s", spew.Sdump(AuthSession)))
	// payload, _, err := jw_token.ValidateToken(AuthSession.JWToken, "")
	// if err != nil {
	// 	return nil, nil, vsLog.Errorf("ValidateToken failed: " + err.Error())
	// }
	// vsLog.Info(fmt.Sprintf("\nPayload: %s\n", spew.Sdump(payload)))
	// vsLog.Info(usr.UserName + " Logged in")

	//fmt.Printf("Login:111  --  User: %s\n", spew.Sdump(usr))
	//loginResponse, err := FillLoginResponse(&usr, token)
	//fmt.Printf("Login:113  --  FillLoginResponse took: %s\n", time.Since(startTime))
	//fmt.Printf("Login:114  --  LoginResponse: %s\n", spew.Sdump(loginResponse))
	//vsLog.Info(fmt.Sprintf("Login Elapsed time: %s", time.Since(loginTime)))
	//fmt.Printf("Login Total:117  --  total Login Elapsed time: %s\n", time.Since(StartTotalTime))
	//return &usr, AuthSession, nil
}

// EncryptPassword: returns an encrypted password for storage or find
func EncryptPassword(passwd string) string {
	return passwd //TODO: Actually encrypt the password

}

func CreateSessionId() string {
	id, _ := nano_uuid.New()
	return id
	// if err!= nil {
	// 	return "", fmt.Errorf("auth_session:91 -- Could not generate uuid: %s\n", err.Error())
	// }
}

// func ValidateSession(id string) (*AuthSession, error) {
// 	return nil, errors.New("Not Implemented")

// }

func CreateToken(ip, userName string, userId, fullName, role string, sessionId string) (string, *jwt.UcPayload, error) {
	//refreshKey := os.Getenv("REFRESH_SECRET")
	// maker, err := jwToken.NewJWTMaker(jwtKey)
	// if err != nil {
	// 	fmt.Printf("NewJWToken err: %s\n", err.Error())
	// 	return "", err
	// }
	token, ucPayload, err := jwt.CreateToken(ip, userName, userId, fullName, role, sessionId)
	if err != nil {
		return "", nil, vsLog.Errorf("CreateToken failed: " + err.Error())
	}
	vsLog.Debug3("Token Payload: " + spew.Sdump(ucPayload))
	return token, ucPayload, err

	// // id, err := nano_uuid.New()  //sessionId
	// // if err != nil {
	// // 	return "", fmt.Errorf("user:133 -- Could not generate uuid: %s\n", err.Error())
	// // }
	// //Creating Session

	// atExpires := time.Now().Add(time.Minute * 60).Unix()
	// td := &TokenDetails{}
	// td.AtExpires = atExpires //time.Now().Add(time.Minute * 30).Unix()  //TODO: Token Expiration should come from config
	// td.AccessUuid, _ = nano_uuid.New()
	// //td.AccessUuid = uuid.NewV4().String()

	// td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix() //TODO: Token Refresh Expiration should come from config
	// //td.RefreshUuid = uuid.NewV4().String()
	// td.RefreshUuid, _ = nano_uuid.New()
	// secretToken := []byte(jwtKey)
	// fmt.Printf("\nSecretToken: %v\n\n", secretToken)
	// atClaims := jwt.MapClaims{}
	// atClaims["authorized"] = true
	// atClaims["user_id"] = userId
	// atClaims["client_id"] = clientId
	// atClaims["exp"] = atExpires
	// //atClaims["session_id"] = id
	// at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	// td.AccessToken, err = at.SignedString([]byte(jwtKey))
	// if err != nil {
	// 	return "", err
	// }
	// //Creating Refresh Token

	// rtClaims := jwt.MapClaims{}
	// rtClaims["refresh_uuid"] = td.RefreshUuid
	// rtClaims["user_id"] = userId
	// rtClaims["client_id"] = clientId
	// rtClaims["exp"] = td.RtExpires
	// rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	// td.RefreshToken, err = rt.SignedString([]byte(refreshKey))
	// if err != nil {
	// 	return "", err
	// }

	// return td.AccessToken, nil
}

// func TokenSignedString(token *jwt.Token) (string, error) {
// 	return token.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
// }

// func VerifyToken(r *http.Request) (*jwt.Token, error) {
// 	tokenString := ExtractToken(r)
// 	return VerifyTokenString(tokenString)
// }

// func VerifyTokenString(tokenString string) (*jwt.Token, error) {
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		//Make sure that the token method conform to "SigningMethodHMAC"
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}
// 		jwtKey := []byte(os.Getenv("ACCESS_SECRET"))
// 		fmt.Printf("accessSecret: %s\n", jwtKey)
// 		return []byte(os.Getenv("ACCESS_SECRET")), nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return token, nil
// }

// func ExtractToken(r *http.Request) string {
// 	bearToken := r.Header.Get("Authorization")
// 	//normally Authorization the_token_xxx
// 	strArr := strings.Split(bearToken, " ")
// 	if len(strArr) == 2 {
// 		vsLog.Info("It is a bearer token")
// 		token := strArr[1]
// 		return token
// 	}
// 	vsLog.Info("It is not a bearer token")
// 	return ""
// }

// func TokenValid(token *jwt.Token) error {
// 	claims, ok := token.Claims.(jwt.Claims)
// 	if !ok && !token.Valid {
// 		return errors.New("token is invalid")
// 	}
// 	vsLog.Debug3("Token>Claim: " + spew.Sdump(claims))
// 	return nil
// }

//   func TokenValid(tkn *jwt.Token) error {
// 	token, err := VerifyToken(r)
// 	if err != nil {
// 	   return err
// 	}
// 	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
// 	   return err
// 	}
// 	return nil
//   }

// func ExtractTokenMetadata(r *http.Request) (*common.AccessDetails, error) {
// 	authToken := r.Header.Get("Authorization")
// 	token, err := VerifyTokenString(authToken)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return GetTokenMetaData(token)
// }

// func GetTokenMetaData(token *jwt.Token) (*common.AccessDetails, error) {
// 	var err error
// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	//if ok && token.Valid {
// 	if ok {
// 		accessUuid, ok := claims["access_uuid"].(string)
// 		if !ok {
// 			return nil, errors.New("access not found in token")
// 		}
// 		sessionId, ok := claims["session_id"].(string)
// 		if !ok {
// 			return nil, errors.New("session not found in token")
// 		}

// 		userId, ok := claims["user_id"].(string)
// 		if !ok {
// 			return nil, errors.New("user not found in token")
// 		}

// 		//    expiresAt, ok := claims["exp"].(int64)
// 		//    if !ok {
// 		// 	  return nil, errors.New("exp not found in token")
// 		//    }

// 		return &common.AccessDetails{
// 			AccessUuid: accessUuid,
// 			SessionId:  sessionId,
// 			UserId:     userId,
// 			Token:      token,
// 			//ExpiresAt: expiresAt,
// 		}, nil
// 	}
// 	return nil, err
// }

// func GetClaimItem(tkn *jwt.Token, claim string) (string, bool) {
// 	claims, ok := tkn.Claims.(jwt.MapClaims)
// 	if !ok {
// 		return "", false
// 	}
// 	return claims[claim].(string), ok
// }

// func SetTokenClaims(tkn *jwt.Token, userId string, customerId string) *jwt.Token {
// 	atClaims := tkn.Claims.(jwt.MapClaims)
// 	//atClaims := jwt.MapClaims{}
// 	atClaims["user_id"] = userId
// 	atClaims["customer_id"] = customerId
// 	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
// 	return at
// 	// newToken, err :=
// 	// td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
// 	// if err != nil {
// 	//    return  nil, err
// 	// }
// 	// vsLog.Debug2f("claims[session_id] = %s", atClaims["session_id"])
// 	// tkn.Claims = atClaims
// 	// //vsLog.Debug2f("Token.claims[session_id] = %s", atClaims["session_id"])
// 	// return tkn
// }

// func UpdateTokenExpire(tkn *jwt.Token) *jwt.Token {
// 	atClaims := tkn.Claims.(jwt.MapClaims)
// 	//atClaims := jwt.MapClaims{}
// 	atClaims["authorized"] = false
// 	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix() //TODO: Token Expiration should come from config
// 	vsLog.Info("UpdTokenExpire: " + spew.Sdump(atClaims))
// 	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
// 	return at
// }

func (user *User) RemoteOrLocal(SystemId string) string {
	for index, element := range user.Locals {
		for idx, system := range element.Systems {
			if system.ID.Hex() == SystemId {
				vsLog.Debug3(fmt.Sprintf("System: %s is a Local system at elements: %d, %d", spew.Sdump(system), index, idx))
				return "local"
			}
		}
	}
	for index, element := range user.Remotes {
		for idx, system := range element.Systems {
			if system.ID.Hex() == SystemId {
				vsLog.Debug3(fmt.Sprintf("System: %s is a Remote system at elements: %d, %d", spew.Sdump(system), index, idx))
				return "remote"
			}
		}
	}
	vsLog.Error(fmt.Sprintf("System: %s is not a valid system id for user: %s", SystemId, user.FullName))
	return "none"
}

func (user *User) Insert() error {
	if user.Exists() {
		return fmt.Errorf("either username or password is invalid")
	}
	tn := time.Now().UTC()
	user.PasswordChangedAt = &tn
	user.CreatedAt = user.PasswordChangedAt
	user.UpdatedAt = user.PasswordChangedAt
	user.Attempts = 0
	user.ID = primitive.NewObjectID()
	err := user.HashPassword(user.Password)
	//hash_pwd, err := HashPassword(user.Password)
	if err != nil {
		//log.Printf("Insert User:376 - error %s", err.Error())
		return vsLog.Errorf("HashPassword failed: " + err.Error())
	}
	//fmt.Printf("Hashed Password: %s\n", user.Password)
	//user.Password = hash_pwd
	collection, _ := GetCollection("user")
	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		return vsLog.Errorf("Insert User failed: " + err.Error())
	}
	return nil
}

func (user *User) Exists() bool {
	u := &common.User{}
	filter := bson.M{"user_name": user.UserName}
	collection, _ := GetCollection("user")
	err := collection.FindOne(context.TODO(), filter).Decode(u)
	return err == nil
}

// func HashPassword(password string) (string, error) {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
// 	if err != nil {
// 		return "", err
// 	}
// 	//user.Password = string(bytes)
// 	//log.Printf("HashPassword:408 - [%s]\n", string(bytes))
// 	return string(bytes), nil
// }

func (user *User) HashPassword(password string) error {
	startTime := time.Now()
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	vsLog.Debug4(fmt.Sprintf("HashPassword took %s", time.Since(startTime)))
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	startTime := time.Now()
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	vsLog.Debug3(fmt.Sprintf("CompareHashAndPassword took %s\n", time.Since(startTime)))
	return nil
}

func (usr *User) Delete() error {
	//startTime := time.Now()
	collection, _ := GetCollection("user")
	filter := bson.M{"_id": usr.ID}
	//vsLog.Debug2f("    bson filter delete: %v\n", filter)
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		vsLog.Info(fmt.Sprintf("User.Delete for userId %s failed: %s", usr.ID.Hex(), err))
		return err
	}
	//vsLog.Infof("@@@!!!   140 -- Deleted %d Users for for ID %s in %s", deleteResult.DeletedCount, usr.ID.Hex(), time.Since(startTime))
	return nil
}

func (user *User) UpdateActive(emr string) error {
	saveUser := *user
	filter := bson.M{"_id": user.ID}
	update := bson.M{}
	tn := time.Now().UTC()
	switch strings.ToLower(emr) {
	case "local":
		user.CurrentLocalPatient.CreatedAt = &tn
		update = bson.M{"$set": bson.M{"currentLocalPatient": user.CurrentLocalPatient}}
	case "remote":
		user.CurrentRemotePatient.CreatedAt = &tn
		update = bson.M{"$set": bson.M{"currentRemotePatient": user.CurrentRemotePatient}}
	}

	collection, err := GetCollection("user")
	updResults, err := collection.UpdateOne(context.TODO(), filter, update)
	vsLog.Debug3("updResult: " + spew.Sdump(updResults))
	if err != nil {
		*user = saveUser
		return vsLog.Errorf("UpdateActive failed: " + err.Error())
	}
	return nil
}

func (user *User) GetActivePatient(destSysId string) (*ActivePatient, error) {

	emrType := user.RemoteOrLocal(destSysId)
	activePatient := &ActivePatient{}
	switch strings.ToLower(emrType) {
	case "local":
		activePatient = user.CurrentLocalPatient
		vsLog.Debug3("LocalActivePatient: " + spew.Sdump(activePatient))
	case "remote":
		activePatient = user.CurrentRemotePatient
		vsLog.Debug3("RemoteActivePatient: " + spew.Sdump(activePatient))
	}
	return activePatient, nil
}

func GetSystemSummary(id primitive.ObjectID) (*common.SystemSummary, error) {
	collection, _ := GetCollection("SystemConfig")
	filter := bson.M{"_id": id}
	//sysCfg := &SystemConfig{}
	sysSum := &common.SystemSummary{}
	// sys := &common.System{}
	// err := collection.FindOne(context.TODO(), filter).Decode(sys)
	// if err != nil {
	// 	return nil, err
	// }
	// vsLog.Info(fmt.Sprintf(" SysConfig: %s", spew.Sdump(sys)))

	err = collection.FindOne(context.TODO(), filter).Decode(sysSum)
	if err != nil {
		return nil, err
	}
	//vsLog.Debug2("GetSystemSummary  --  SysSum: " + spew.Sdump(sysSum))
	return sysSum, nil //vsLog.Errorf("SysSum: " + spew.Sdump(sysSum))
}

func GetSystem(id primitive.ObjectID) (*common.System, error) {
	collection, _ := GetCollection("SystemConfig")
	filter := bson.M{"_id": id}
	sysCfg := common.System{}
	//sysCfg := &SystemConfig{}
	sys := &common.System{}
	err := collection.FindOne(context.TODO(), filter).Decode(sysCfg)
	if err != nil {
		return nil, err
	}
	vsLog.Debug2("SysConfig: " + spew.Sdump(sysCfg))
	sys.ID = sysCfg.ID
	//sys.ConnectorUrl = sysCfg.ConnectorConfig.URL
	sys.DisplayName = sysCfg.DisplayName
	///sys.FacilityId = sysCfg.FacilityId
	sys.Name = sysCfg.Name
	return sys, nil
}

func GetFacility(id primitive.ObjectID) (*common.Facility, error) {
	var systems []*common.SystemSummary
	collection, _ := GetCollection("facilities")
	filter := bson.M{"_id": id}
	sysFac := &common.Facility{}
	fac := &common.Facility{}
	err := collection.FindOne(context.TODO(), filter).Decode(sysFac)
	if err != nil {
		return nil, err
	}
	vsLog.Debug3("SysFacility: " + spew.Sdump(sysFac))
	fac.ID = sysFac.ID
	fac.Description = sysFac.Description
	fac.DisplayName = sysFac.DisplayName
	fac.Name = sysFac.Name
	//fac.IconLogo = sysFac.IconLogo
	// fac.Lat = sysFac.Lat
	// fac.Lon = sysFac.Lon

	for _, s := range sysFac.Systems {
		sys, err := GetSystemSummary(s.ID)
		if err != nil {
			vsLog.Errorf(fmt.Sprintf("System: %s not found: err: %s", s.ID, err.Error()))
			continue
		}
		vsLog.Debug3("Adding Facility:" + spew.Sdump(sys))
		systems = append(systems, sys)
	}
	fac.Systems = systems

	return fac, nil
}

func FillLoginResponse(usr *User, token string) (*LoginResponse, error) {
	//vsLog.Debug3("FillLoginResponse" + spew.Sdump(usr))
	lr := LoginResponse{}
	lr.FullName = usr.FullName
	lr.ID = usr.ID
	lr.Message = "No Message"
	lr.Role = usr.Role
	lr.Token = token
	lr.UserName = usr.UserName
	lr.Remotes = usr.Remotes
	//vsLog.Debug2("FillLoginResponse- Remotes" + spew.Sdump(lr.Remotes))
	lr.Locals = usr.Locals
	lr.CurrentLocalPatient = usr.CurrentLocalPatient
	lr.CurrentRemotePatient = usr.CurrentRemotePatient
	vsLog.Debug3("FillLoginResponse- Remotes" + spew.Sdump(lr.Remotes))
	vsLog.Debug2("FillLoginResponse- Locals" + spew.Sdump(lr.Locals))
	return &lr, nil
}

func GetUserByUserName(userName string) (*User, error) {
	startTime := time.Now()
	filter := bson.M{"user_name": userName}
	collection, _ := GetCollection("user")
	usr := &User{}
	vsLog.Debug3(fmt.Sprintf("Calling FindOne User: %v", filter))
	err := collection.FindOne(context.TODO(), filter).Decode(usr)
	vsLog.Debug3(fmt.Sprintf("FindOneUser took %s", time.Since(startTime)))
	return usr, err

}

func GetUserById(uid string) (*User, error) {
	//startTime := time.Now()
	userId, err := primitive.ObjectIDFromHex(uid)
	if err != nil {
		return nil, vsLog.Errorf("userId: " + uid + " is invalid")
	}
	filter := bson.M{"_id": userId}
	collection, _ := GetCollection("user")
	usr := &User{}
	vsLog.Debug3(fmt.Sprintf("Calling FindOne User: %v", filter))
	err = collection.FindOne(context.TODO(), filter).Decode(usr)
	//vsLog.Debug3(fmt.Sprintf("FindOneUser took %s", time.Since(startTime)))
	return usr, err
}

func SystemsForFacility(sysFac *common.Facility) ([]*common.System, error) {
	systems := []*common.System{}
	for _, sys := range sysFac.Systems {
		sys, err := GetSystem(sys.ID)
		if err != nil {
			vsLog.Errorf(fmt.Sprintf("System: %s not found: err: %s", sys.ID, err.Error()))
			continue
		}
		vsLog.Debug3("Adding Facility: " + spew.Sdump(sys))
		systems = append(systems, sys)
	}
	vsLog.Debug3("Systems = " + spew.Sdump(systems))
	return systems, nil
}

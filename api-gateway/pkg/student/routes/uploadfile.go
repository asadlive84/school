package routes

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	pb "github.com/asadlive84/school/pb/student"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/timestamppb"
	// "github.com/golang/protobuf/ptypes/timestamp"
)

func UploadFile(ctx *gin.Context, d pb.StudentServiceClient) {
	// Retrieve the uploaded file from the request

	sessionYear := ctx.PostForm("sessionYear")

	file, _, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}
	defer file.Close()

	// Create a temporary file to store the uploaded file
	tempFile, err := os.CreateTemp("", "uploaded-*.csv")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create temporary file"})
		return
	}
	defer os.Remove(tempFile.Name())

	// Copy the uploaded file to the temporary file
	_, err = io.Copy(tempFile, file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write file to disk"})
		return
	}

	// Open the temporary file for reading
	csvFile, err := os.Open(tempFile.Name())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open temporary file"})
		return
	}
	defer csvFile.Close()

	// Create a CSV reader
	reader := csv.NewReader(csvFile)

	// Read the first record to get the header names
	headers, err := reader.Read()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read header"})
		return
	}

	// Read the remaining records as values
	var values [][]string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read CSV file"})
			return
		}
		values = append(values, record)
	}

	// Process the headers and values as needed
	processData(sessionYear, headers, values, d)

	ctx.JSON(http.StatusOK, gin.H{"message": "File uploaded and processed successfully"})

}

func processData(sessionYear string, headers []string, values [][]string, d pb.StudentServiceClient) {
	// Process the headers (column names)
	for _, header := range headers {
		// Process each header as needed
		// For example, print the header
		println("Header:", header)
	}

	// Process the values (rows)
	for _, row := range values {
		// Process each row of values
		dateString := row[5]
		var protoTimestamp *timestamppb.Timestamp

		if dateString != "" {
			// Parse the input string into a time.Time value
			layout := "02/01/2006" // "02" is for day, "01" is for month, and "2006" is for year
			parsedTime, err := time.Parse(layout, dateString)
			if err != nil {
				fmt.Println("Error parsing date:", err)
			}
			protoTimestamp, err = ptypes.TimestampProto(parsedTime)
			if err != nil {
				fmt.Println("Error converting to protobuf timestamp:", err)
				// return
			}
		}

		gender := "MALE"
		if strings.ToLower(row[6]) == "female" {
			gender = "FEMALE"
		}

		religion := "ISLAM"
		if strings.ToLower(row[15]) != "islam" {
			religion = "HINDU"
		}

		if len(religion) >= 10 {
			religion = "N/A"
		}

		addressInfo := strings.Split(row[14], ",")

		district := strings.TrimRight(strings.TrimLeft(addressInfo[3], " "), " ")
		district = strings.ReplaceAll(district, " ", "")
		district = strings.ReplaceAll(district, ".", "")

		upzilla := strings.TrimRight(strings.TrimLeft(addressInfo[2], " "), " ")
		upzilla = strings.ReplaceAll(upzilla, " ", "")
		upzilla = strings.ReplaceAll(upzilla, ".", "")

		unionName := strings.TrimRight(strings.TrimLeft(addressInfo[1], " "), " ")
		unionName = strings.ReplaceAll(unionName, " ", "")
		unionName = strings.ReplaceAll(unionName, ".", "")

		villageOrRoad := strings.TrimRight(strings.TrimLeft(addressInfo[0], " "), " ")
		villageOrRoad = strings.ReplaceAll(villageOrRoad, " ", "")
		villageOrRoad = strings.ReplaceAll(villageOrRoad, ".", "")

		std := &pb.Student{
			Id:            "",
			StdId:         strings.TrimRight(strings.TrimLeft(row[0], " "), " "),
			Name:          strings.TrimRight(strings.TrimLeft(row[1], " "), " "),
			NameBn:        strings.TrimRight(strings.TrimLeft(row[2], " "), " "),
			FathersName:   strings.TrimRight(strings.TrimLeft(row[3], " "), " "),
			MothersName:   strings.TrimRight(strings.TrimLeft(row[4], " "), " "),
			Dob:           protoTimestamp,
			Gender:        gender,
			BloodGroup:    strings.TrimRight(strings.TrimLeft(row[7], " "), " "),
			MobileNumber:  strings.TrimRight(strings.TrimLeft(row[8], " "), " "),
			Session:       strings.TrimRight(strings.TrimLeft(row[9], " "), " "),
			ClassName:     strings.TrimRight(strings.TrimLeft(row[10], " "), " "),
			ClassSection:  "",
			ClassGroup:    "",
			ClassRoll:     "",
			Address:       strings.TrimRight(strings.TrimLeft(row[14], " "), " "),
			Religion:      religion,
			District:      strings.ToLower(district),
			Upzilla:       strings.ToLower(upzilla),
			UnionName:     strings.ToLower(unionName),
			VillageOrRoad: strings.ToLower(villageOrRoad),
			
		}

		ctx := context.Background()
		stdResp, err := d.Insert(ctx, &pb.InsertRequest{
			Student: std,
		})
		if err != nil {
			fmt.Println("Error getting stdResp:", err)
			continue
		}

		sesYr, err := strconv.Atoi(sessionYear)
		if err != nil {
			fmt.Println("Error getting sessionYear:", err)
			return
		}

		d.InsertAcademicSession(ctx, &pb.InsertAcademicSessionRequest{
			StudentId: stdResp.StdId,
			SessionId: int32(sesYr),
		})

	}
}

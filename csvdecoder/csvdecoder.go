package csvdecoder

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

type Person struct {
    Key 							int `json:"key"`
    Id     							int `json:"id"`
    Uid      						string    `json:"uid"`
	Domain      					string    `json:"domain"`
	Cn      						string    `json:"cn"`
	Department      				string    `json:"department"`
	Title      						string    `json:"title"`
	Who      						string    `json:"who"`
	Logon_count     				int    `json:"logon_count"`
	Num_logons7     				int    `json:"Num_logons7"`
	Num_share7						int    `json:"Num_share7"`
	Num_file7	 					int    `json:"num_file7"`
	Num_ad7	 						int    `json:"num_ad7"`
	Num_n7	 		int    `json:"num_n7"`
	Num_logons14	int    `json:"num_logons14"`
	Num_share14 	int    `json:"num_share14"`
	Num_file14		int    `json:"num_file14"`
	Num_ad14		int    `json:"num_ad14"`
	Num_n14			int    `json:"num_n14"`
	Num_logons30 	int    `json:"num_logons30"`
	Num_share30 	int    `json:"num_share30"`
	Num_file30 		int    `json:"num_file30"`
	Num_ad30 		int    `json:"num_ad30"`
	Num_n30 		int    `json:"num_n30"`
	Num_logons150 	int    `json:"num_logons150"`
	Num_share150 int    `json:"num_share150"`
	Num_file150 int    `json:"num_file150"`
	Num_ad150 int    `json:"num_ad150"`
	Num_n150 int    `json:"num_n150"`
	Num_logons365 int    `json:"num_logons365"`
	Num_share365 int    `json:"num_share365"`
	Num_file365 int    `json:"num_file365"`
	Num_ad365 int    `json:"num_ad365"`
	Num_n365 int    `json:"num_n365"`
	Has_user_principal_name int    `json:"has_user_principal_name"`
	Has_mail int    `json:"has_mail"`
	Has_phone int    `json:"has_phone"`
	Flag_disabled int    `json:"flag_disabled"`
	Flag_lockout int    `json:"flag_lockout"`
	Flag_password_not_required int    `json:"flag_password_not_required"`
	Flag_password_cant_change int    `json:"flag_password_cant_change"`
	Flag_dont_expire_password int    `json:"flag_dont_expire_password"`
	Owned_files int    `json:"owned_files"`
	Num_mailboxes int    `json:"num_mailboxes"`
	Num_member_of_groups int    `json:"num_member_of_groups"`
	Num_member_of_indirect_groups int    `json:"num_member_of_indirect_groups"`
	Member_of_indirect_groups_ids int    `json:"member_of_indirect_groups_ids"`
	Member_of_groups_ids int    `json:"member_of_groups_ids"`
	Is_admin int    `json:"is_admin"`
	Is_service int    `json:"is_service"`
}
 
func Convert(nameCSV string) (persons []Person){
	csvFile, err := os.Open(nameCSV)
	if err != nil {
		log.Println(err)
	}
	defer csvFile.Close()
   
	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1
   
	csvData, err := reader.ReadAll()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
   
	var per Person
   
	for _, each := range csvData {
		per.Key,_ = strconv.Atoi(each[0])
		per.Id, _ = strconv.Atoi(each[1])
		per.Uid = each[2]
		per.Domain = each[3]
		per.Cn = each[4]
		per.Department = each[5]
		per.Title = each[6]
		per.Who = each[7]
		per.Logon_count, _  = strconv.Atoi(each[8])

		per.Num_logons7, _  = strconv.Atoi(each[9])
		per.Num_share7, _  = strconv.Atoi(each[10])
		per.Num_file7, _  = strconv.Atoi(each[11])
		per.Num_ad7, _  = strconv.Atoi(each[12])
		per.Num_n7, _  = strconv.Atoi(each[13])

		per.Num_logons14, _  = strconv.Atoi(each[14])
		per.Num_share14, _  = strconv.Atoi(each[15])
		per.Num_file14, _  = strconv.Atoi(each[16])
		per.Num_ad14, _  = strconv.Atoi(each[17])
		per.Num_n14, _  = strconv.Atoi(each[18])

		per.Num_logons30, _  = strconv.Atoi(each[19])
		per.Num_share30, _  = strconv.Atoi(each[20])
		per.Num_file30, _  = strconv.Atoi(each[21])
		per.Num_ad30, _  = strconv.Atoi(each[22])
		per.Num_n30, _  = strconv.Atoi(each[23])

		per.Num_logons150, _  = strconv.Atoi(each[24])
		per.Num_share150, _  = strconv.Atoi(each[25])
		per.Num_file150, _  = strconv.Atoi(each[26])
		per.Num_ad150, _  = strconv.Atoi(each[27])
		per.Num_n150, _  = strconv.Atoi(each[28])

		per.Num_logons365, _  = strconv.Atoi(each[29])
		per.Num_share365, _  = strconv.Atoi(each[30])
		per.Num_file365, _  = strconv.Atoi(each[31])
		per.Num_ad365, _  = strconv.Atoi(each[32])
		per.Num_n365, _  = strconv.Atoi(each[33])

		per.Has_user_principal_name, _  = strconv.Atoi(each[34])
		per.Has_mail, _  = strconv.Atoi(each[35])
		per.Has_phone, _  = strconv.Atoi(each[36])

		per.Flag_disabled, _  = strconv.Atoi(each[37])
		per.Flag_lockout, _  = strconv.Atoi(each[38])
		per.Flag_password_not_required, _  = strconv.Atoi(each[39])
		per.Flag_password_cant_change, _  = strconv.Atoi(each[40])
		per.Flag_dont_expire_password, _  = strconv.Atoi(each[41])

		per.Owned_files, _  = strconv.Atoi(each[42])
		per.Num_mailboxes, _  = strconv.Atoi(each[43])
		per.Num_member_of_groups, _  = strconv.Atoi(each[44])
		per.Num_member_of_indirect_groups, _  = strconv.Atoi(each[45])

		per.Member_of_indirect_groups_ids, _  = strconv.Atoi(each[46])
		per.Num_member_of_groups, _  = strconv.Atoi(each[47])

		per.Is_admin, _  = strconv.Atoi(each[48])
		per.Is_service, _  = strconv.Atoi(each[49])

		persons = append(persons, per)
		
	}
	return
}
   
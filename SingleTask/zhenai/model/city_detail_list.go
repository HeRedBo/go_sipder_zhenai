package model


type CityDetailList struct {
	MemberListData struct {
		CurrentPage int `json:"currentPage"`
		MemberList  []struct {
			Age              int    `json:"age"`
			AvatarURL        string `json:"avatarURL"`
			Car              string `json:"car"`
			Children         string `json:"children"`
			Constellation    string `json:"constellation"`
			Education        string `json:"education"`
			Height           int    `json:"height"`
			House            string `json:"house"`
			IntroduceContent string `json:"introduceContent"`
			IsRecommend      int    `json:"isRecommend"`
			LastModTime      string `json:"lastModTime"`
			Marriage         string `json:"marriage"`
			MemberID         int    `json:"memberID"`
			NickName         string `json:"nickName"`
			ObjectAge        string `json:"objectAge"`
			ObjectHight      string `json:"objectHight"`
			ObjectMarriage   string `json:"objectMarriage"`
			ObjectSalary     string `json:"objectSalary"`
			Occupation       string `json:"occupation"`
			Salary           string `json:"salary"`
			Sex              int    `json:"sex"`
			WorkCity         string `json:"workCity"`
		} `json:"memberList"`
		PageInfos []struct {
			CurrLocation bool   `json:"currLocation"`
			Link         bool   `json:"link"`
			PageContent  string `json:"pageContent"`
			PageLink     string `json:"pageLink,omitempty"`
		} `json:"pageInfos"`
		TkdInfo struct {
			Desc     string `json:"desc"`
			Keywords string `json:"keywords"`
			Title    string `json:"title"`
			URL      string `json:"url"`
		} `json:"tkdInfo"`
		Total int `json:"total"`
	} `json:"memberListData"`
	RecommendListData struct {
		MemberList []struct {
			Age                   int    `json:"age"`
			AvatarURL             string `json:"avatarURL"`
			BriefIntroduceContent string `json:"briefIntroduceContent"`
			Height                int    `json:"height"`
			IntroduceContent      string `json:"introduceContent"`
			LinkURL               string `json:"linkURL"`
			MemberID              int    `json:"memberId"`
			Nickname              string `json:"nickname"`
			Salary                int    `json:"salary"`
			Sex                   int    `json:"sex"`
		} `json:"memberList"`
	} `json:"recommendListData"`
	FooterData struct {
		DistrictLevel int `json:"districtLevel"`
		MainCategory  struct {
			Desc         string `json:"desc"`
			MainCategory []struct {
				Content   string `json:"content"`
				IslinkURL bool   `json:"islinkURL"`
				LinkURL   string `json:"linkURL,omitempty"`
			} `json:"mainCategory"`
		} `json:"mainCategory"`
		NearbyCity struct {
			CityList []struct {
				LinkContent string `json:"linkContent"`
				LinkURL     string `json:"linkURL"`
			} `json:"cityList"`
			Desc string `json:"desc"`
		} `json:"nearbyCity"`
		RelateCategory struct {
			Desc             string `json:"desc"`
			RelatedCategorys []struct {
				CategoryWord string `json:"categoryWord"`
				LinkURL      string `json:"linkURL"`
			} `json:"relatedCategorys"`
		} `json:"relateCategory"`
	} `json:"footerData"`
	NavigationData struct {
		District []struct {
			CurrLocation bool   `json:"currLocation"`
			LinkContent  string `json:"linkContent"`
			LinkURL      string `json:"linkURL"`
		} `json:"district"`
		Location []struct {
			CurrLocation bool   `json:"currLocation"`
			LinkContent  string `json:"linkContent"`
			LinkURL      string `json:"linkURL"`
		} `json:"location"`
		Sex []struct {
			CurrLocation bool   `json:"currLocation"`
			LinkContent  string `json:"linkContent"`
			LinkURL      string `json:"linkURL"`
		} `json:"sex"`
	} `json:"navigationData"`
	RecommendRegisterListData struct {
		MemberList []struct {
			Age                   int    `json:"age"`
			AvatarURL             string `json:"avatarURL"`
			BriefIntroduceContent string `json:"briefIntroduceContent"`
			Height                int    `json:"height"`
			IntroduceContent      string `json:"introduceContent"`
			LinkURL               string `json:"linkURL"`
			MemberID              int    `json:"memberId"`
			Nickname              string `json:"nickname"`
			Salary                int    `json:"salary"`
			Sex                   int    `json:"sex"`
		} `json:"memberList"`
	} `json:"recommendRegisterListData"`
}
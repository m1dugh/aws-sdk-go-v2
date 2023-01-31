// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type Engine string

// Enum values for Engine
const (
	EngineStandard Engine = "standard"
	EngineNeural   Engine = "neural"
)

// Values returns all known values for Engine. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Engine) Values() []Engine {
	return []Engine{
		"standard",
		"neural",
	}
}

type Gender string

// Enum values for Gender
const (
	GenderFemale Gender = "Female"
	GenderMale   Gender = "Male"
)

// Values returns all known values for Gender. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Gender) Values() []Gender {
	return []Gender{
		"Female",
		"Male",
	}
}

type LanguageCode string

// Enum values for LanguageCode
const (
	LanguageCodeArb     LanguageCode = "arb"
	LanguageCodeCmnCn   LanguageCode = "cmn-CN"
	LanguageCodeCyGb    LanguageCode = "cy-GB"
	LanguageCodeDaDk    LanguageCode = "da-DK"
	LanguageCodeDeDe    LanguageCode = "de-DE"
	LanguageCodeEnAu    LanguageCode = "en-AU"
	LanguageCodeEnGb    LanguageCode = "en-GB"
	LanguageCodeEnGbWls LanguageCode = "en-GB-WLS"
	LanguageCodeEnIn    LanguageCode = "en-IN"
	LanguageCodeEnUs    LanguageCode = "en-US"
	LanguageCodeEsEs    LanguageCode = "es-ES"
	LanguageCodeEsMx    LanguageCode = "es-MX"
	LanguageCodeEsUs    LanguageCode = "es-US"
	LanguageCodeFrCa    LanguageCode = "fr-CA"
	LanguageCodeFrFr    LanguageCode = "fr-FR"
	LanguageCodeIsIs    LanguageCode = "is-IS"
	LanguageCodeItIt    LanguageCode = "it-IT"
	LanguageCodeJaJp    LanguageCode = "ja-JP"
	LanguageCodeHiIn    LanguageCode = "hi-IN"
	LanguageCodeKoKr    LanguageCode = "ko-KR"
	LanguageCodeNbNo    LanguageCode = "nb-NO"
	LanguageCodeNlNl    LanguageCode = "nl-NL"
	LanguageCodePlPl    LanguageCode = "pl-PL"
	LanguageCodePtBr    LanguageCode = "pt-BR"
	LanguageCodePtPt    LanguageCode = "pt-PT"
	LanguageCodeRoRo    LanguageCode = "ro-RO"
	LanguageCodeRuRu    LanguageCode = "ru-RU"
	LanguageCodeSvSe    LanguageCode = "sv-SE"
	LanguageCodeTrTr    LanguageCode = "tr-TR"
	LanguageCodeEnNz    LanguageCode = "en-NZ"
	LanguageCodeEnZa    LanguageCode = "en-ZA"
	LanguageCodeCaEs    LanguageCode = "ca-ES"
	LanguageCodeDeAt    LanguageCode = "de-AT"
	LanguageCodeYueCn   LanguageCode = "yue-CN"
	LanguageCodeArAe    LanguageCode = "ar-AE"
	LanguageCodeFiFi    LanguageCode = "fi-FI"
)

// Values returns all known values for LanguageCode. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (LanguageCode) Values() []LanguageCode {
	return []LanguageCode{
		"arb",
		"cmn-CN",
		"cy-GB",
		"da-DK",
		"de-DE",
		"en-AU",
		"en-GB",
		"en-GB-WLS",
		"en-IN",
		"en-US",
		"es-ES",
		"es-MX",
		"es-US",
		"fr-CA",
		"fr-FR",
		"is-IS",
		"it-IT",
		"ja-JP",
		"hi-IN",
		"ko-KR",
		"nb-NO",
		"nl-NL",
		"pl-PL",
		"pt-BR",
		"pt-PT",
		"ro-RO",
		"ru-RU",
		"sv-SE",
		"tr-TR",
		"en-NZ",
		"en-ZA",
		"ca-ES",
		"de-AT",
		"yue-CN",
		"ar-AE",
		"fi-FI",
	}
}

type OutputFormat string

// Enum values for OutputFormat
const (
	OutputFormatJson      OutputFormat = "json"
	OutputFormatMp3       OutputFormat = "mp3"
	OutputFormatOggVorbis OutputFormat = "ogg_vorbis"
	OutputFormatPcm       OutputFormat = "pcm"
)

// Values returns all known values for OutputFormat. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (OutputFormat) Values() []OutputFormat {
	return []OutputFormat{
		"json",
		"mp3",
		"ogg_vorbis",
		"pcm",
	}
}

type SpeechMarkType string

// Enum values for SpeechMarkType
const (
	SpeechMarkTypeSentence SpeechMarkType = "sentence"
	SpeechMarkTypeSsml     SpeechMarkType = "ssml"
	SpeechMarkTypeViseme   SpeechMarkType = "viseme"
	SpeechMarkTypeWord     SpeechMarkType = "word"
)

// Values returns all known values for SpeechMarkType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SpeechMarkType) Values() []SpeechMarkType {
	return []SpeechMarkType{
		"sentence",
		"ssml",
		"viseme",
		"word",
	}
}

type TaskStatus string

// Enum values for TaskStatus
const (
	TaskStatusScheduled  TaskStatus = "scheduled"
	TaskStatusInProgress TaskStatus = "inProgress"
	TaskStatusCompleted  TaskStatus = "completed"
	TaskStatusFailed     TaskStatus = "failed"
)

// Values returns all known values for TaskStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (TaskStatus) Values() []TaskStatus {
	return []TaskStatus{
		"scheduled",
		"inProgress",
		"completed",
		"failed",
	}
}

type TextType string

// Enum values for TextType
const (
	TextTypeSsml TextType = "ssml"
	TextTypeText TextType = "text"
)

// Values returns all known values for TextType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (TextType) Values() []TextType {
	return []TextType{
		"ssml",
		"text",
	}
}

type VoiceId string

// Enum values for VoiceId
const (
	VoiceIdAditi     VoiceId = "Aditi"
	VoiceIdAmy       VoiceId = "Amy"
	VoiceIdAstrid    VoiceId = "Astrid"
	VoiceIdBianca    VoiceId = "Bianca"
	VoiceIdBrian     VoiceId = "Brian"
	VoiceIdCamila    VoiceId = "Camila"
	VoiceIdCarla     VoiceId = "Carla"
	VoiceIdCarmen    VoiceId = "Carmen"
	VoiceIdCeline    VoiceId = "Celine"
	VoiceIdChantal   VoiceId = "Chantal"
	VoiceIdConchita  VoiceId = "Conchita"
	VoiceIdCristiano VoiceId = "Cristiano"
	VoiceIdDora      VoiceId = "Dora"
	VoiceIdEmma      VoiceId = "Emma"
	VoiceIdEnrique   VoiceId = "Enrique"
	VoiceIdEwa       VoiceId = "Ewa"
	VoiceIdFiliz     VoiceId = "Filiz"
	VoiceIdGabrielle VoiceId = "Gabrielle"
	VoiceIdGeraint   VoiceId = "Geraint"
	VoiceIdGiorgio   VoiceId = "Giorgio"
	VoiceIdGwyneth   VoiceId = "Gwyneth"
	VoiceIdHans      VoiceId = "Hans"
	VoiceIdInes      VoiceId = "Ines"
	VoiceIdIvy       VoiceId = "Ivy"
	VoiceIdJacek     VoiceId = "Jacek"
	VoiceIdJan       VoiceId = "Jan"
	VoiceIdJoanna    VoiceId = "Joanna"
	VoiceIdJoey      VoiceId = "Joey"
	VoiceIdJustin    VoiceId = "Justin"
	VoiceIdKarl      VoiceId = "Karl"
	VoiceIdKendra    VoiceId = "Kendra"
	VoiceIdKevin     VoiceId = "Kevin"
	VoiceIdKimberly  VoiceId = "Kimberly"
	VoiceIdLea       VoiceId = "Lea"
	VoiceIdLiv       VoiceId = "Liv"
	VoiceIdLotte     VoiceId = "Lotte"
	VoiceIdLucia     VoiceId = "Lucia"
	VoiceIdLupe      VoiceId = "Lupe"
	VoiceIdMads      VoiceId = "Mads"
	VoiceIdMaja      VoiceId = "Maja"
	VoiceIdMarlene   VoiceId = "Marlene"
	VoiceIdMathieu   VoiceId = "Mathieu"
	VoiceIdMatthew   VoiceId = "Matthew"
	VoiceIdMaxim     VoiceId = "Maxim"
	VoiceIdMia       VoiceId = "Mia"
	VoiceIdMiguel    VoiceId = "Miguel"
	VoiceIdMizuki    VoiceId = "Mizuki"
	VoiceIdNaja      VoiceId = "Naja"
	VoiceIdNicole    VoiceId = "Nicole"
	VoiceIdOlivia    VoiceId = "Olivia"
	VoiceIdPenelope  VoiceId = "Penelope"
	VoiceIdRaveena   VoiceId = "Raveena"
	VoiceIdRicardo   VoiceId = "Ricardo"
	VoiceIdRuben     VoiceId = "Ruben"
	VoiceIdRussell   VoiceId = "Russell"
	VoiceIdSalli     VoiceId = "Salli"
	VoiceIdSeoyeon   VoiceId = "Seoyeon"
	VoiceIdTakumi    VoiceId = "Takumi"
	VoiceIdTatyana   VoiceId = "Tatyana"
	VoiceIdVicki     VoiceId = "Vicki"
	VoiceIdVitoria   VoiceId = "Vitoria"
	VoiceIdZeina     VoiceId = "Zeina"
	VoiceIdZhiyu     VoiceId = "Zhiyu"
	VoiceIdAria      VoiceId = "Aria"
	VoiceIdAyanda    VoiceId = "Ayanda"
	VoiceIdArlet     VoiceId = "Arlet"
	VoiceIdHannah    VoiceId = "Hannah"
	VoiceIdArthur    VoiceId = "Arthur"
	VoiceIdDaniel    VoiceId = "Daniel"
	VoiceIdLiam      VoiceId = "Liam"
	VoiceIdPedro     VoiceId = "Pedro"
	VoiceIdKajal     VoiceId = "Kajal"
	VoiceIdHiujin    VoiceId = "Hiujin"
	VoiceIdLaura     VoiceId = "Laura"
	VoiceIdElin      VoiceId = "Elin"
	VoiceIdIda       VoiceId = "Ida"
	VoiceIdSuvi      VoiceId = "Suvi"
	VoiceIdOla       VoiceId = "Ola"
	VoiceIdHala      VoiceId = "Hala"
	VoiceIdAndres    VoiceId = "Andres"
	VoiceIdSergio    VoiceId = "Sergio"
	VoiceIdRemi      VoiceId = "Remi"
	VoiceIdAdriano   VoiceId = "Adriano"
	VoiceIdThiago    VoiceId = "Thiago"
	VoiceIdRuth      VoiceId = "Ruth"
	VoiceIdStephen   VoiceId = "Stephen"
)

// Values returns all known values for VoiceId. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (VoiceId) Values() []VoiceId {
	return []VoiceId{
		"Aditi",
		"Amy",
		"Astrid",
		"Bianca",
		"Brian",
		"Camila",
		"Carla",
		"Carmen",
		"Celine",
		"Chantal",
		"Conchita",
		"Cristiano",
		"Dora",
		"Emma",
		"Enrique",
		"Ewa",
		"Filiz",
		"Gabrielle",
		"Geraint",
		"Giorgio",
		"Gwyneth",
		"Hans",
		"Ines",
		"Ivy",
		"Jacek",
		"Jan",
		"Joanna",
		"Joey",
		"Justin",
		"Karl",
		"Kendra",
		"Kevin",
		"Kimberly",
		"Lea",
		"Liv",
		"Lotte",
		"Lucia",
		"Lupe",
		"Mads",
		"Maja",
		"Marlene",
		"Mathieu",
		"Matthew",
		"Maxim",
		"Mia",
		"Miguel",
		"Mizuki",
		"Naja",
		"Nicole",
		"Olivia",
		"Penelope",
		"Raveena",
		"Ricardo",
		"Ruben",
		"Russell",
		"Salli",
		"Seoyeon",
		"Takumi",
		"Tatyana",
		"Vicki",
		"Vitoria",
		"Zeina",
		"Zhiyu",
		"Aria",
		"Ayanda",
		"Arlet",
		"Hannah",
		"Arthur",
		"Daniel",
		"Liam",
		"Pedro",
		"Kajal",
		"Hiujin",
		"Laura",
		"Elin",
		"Ida",
		"Suvi",
		"Ola",
		"Hala",
		"Andres",
		"Sergio",
		"Remi",
		"Adriano",
		"Thiago",
		"Ruth",
		"Stephen",
	}
}

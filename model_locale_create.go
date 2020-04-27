package phrase
// LocaleCreate struct for LocaleCreate
type LocaleCreate struct {
	// specify the branch to use
	Branch string `json:"branch,omitempty"`
	// Locale name
	Name string `json:"name,omitempty"`
	// Locale ISO code
	Code string `json:"code,omitempty"`
	// Indicates whether locale is the default locale. If set to true, the previous default locale the project is no longer the default locale.
	Default string `json:"default,omitempty"`
	// Indicates whether locale is a main locale. Main locales are part of the <a href=\"https://help.phrase.com/help/verification-and-proofreading\" target=\"_blank\">Verification System</a> feature.
	Main string `json:"main,omitempty"`
	// Indicates whether locale is a RTL (Right-to-Left) locale.
	Rtl string `json:"rtl,omitempty"`
	// Source locale. Can be the name or public id of the locale. Preferred is the public id.
	SourceLocaleId string `json:"source_locale_id,omitempty"`
	// Indicates that new translations for this locale should be marked as unverified. Part of the <a href=\"https://help.phrase.com/help/verification-and-proofreading\" target=\"_blank\">Advanced Workflows</a> feature.
	UnverifyNewTranslations string `json:"unverify_new_translations,omitempty"`
	// Indicates that updated translations for this locale should be marked as unverified. Part of the <a href=\"https://help.phrase.com/help/verification-and-proofreading\" target=\"_blank\">Advanced Workflows</a> feature.
	UnverifyUpdatedTranslations string `json:"unverify_updated_translations,omitempty"`
	// If set, translations for this locale will be fetched automatically, right after creation.
	Autotranslate string `json:"autotranslate,omitempty"`
}

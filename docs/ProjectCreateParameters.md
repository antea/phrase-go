# ProjectCreateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the project | [optional] 
**MainFormat** | **string** | Main file format specified by its API Extension name. Used for locale downloads if no format is specified. For API Extension names of available file formats see &lt;a href&#x3D;\&quot;https://help.phrase.com/help/supported-platforms-and-formats\&quot;&gt;Format Guide&lt;/a&gt; or our &lt;a href&#x3D;\&quot;#formats\&quot;&gt;Formats API Endpoint&lt;/a&gt;. | [optional] 
**SharesTranslationMemory** | **bool** | Indicates whether the project should share the account&#39;s translation memory | [optional] 
**ProjectImage** | [***os.File**](*os.File.md) | Image to identify the project | [optional] 
**RemoveProjectImage** | **bool** | Indicates whether the project image should be deleted. | [optional] 
**AccountId** | **string** | Account ID to specify the actual account the project should be created in. Required if the requesting user is a member of multiple accounts. | [optional] 
**SourceProjectId** | **string** | When a source project ID is given, a clone of that project will be created, including all locales, keys and translations as well as the main project settings if they are not defined otherwise through the params. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



/*
 * Secure Access Cloud API
 *
 *  ## Introduction  Secure Access Cloud API uses common RESTful resourced based URL conventions and JSON as the exchange format. <br> Properties names are case-sensitive. <br> Some of Secure Access Cloud API calls omit None values from the API response.  The base-URL is `api.`&lt;`tenant-name`&gt;`.luminatesec.com`. For example, if your administration portal URL is _admin.acme.luminatesec.com_, then your API base-URL is _api.acme.luminatesec.com_.  All examples below are performed on a tenant called acme.  ## Common Operations Steps  Below you may find a list of common operations and the relevant API calls for each. Each of these operations can also be performed by using the administrative portal at https://admin.acme.luminatesec.com.  <ol>   <li>     Creating a site and deploying a connector:     <ol type=\"a\">       <li>Creating a new site using the <a href=\"#operation/createSite\">Create site API</a>.<br></li>       <li>         Once a site is created you can use its Id (returned in the response of the Create Site request)         and call the <a href=\"#operation/createConnector\">Create connector API</a>. <br>       </li>       <li>         Deploy the Secure Access Cloud connector:         <ol type=\"i\">           <li>Retrieve the deployment command using the <a href=\"#operation/getConnectorCommand\">Connector Deployment Command API.</a> <br> </li>           <li>Execute the command on the target machine.</li>         </ol>       </li>     </ol>   </li>   <li>     Creating an application:       <ol type=\"a\">         <li>           An application is always associated with a specific site for routing the traffic to the application via the connectors associated with the same site.           In order to create the application, call the <a href=\"#operation/createApplication\">Create Application API</a>         </li>         <li>           Once the application is created, you *must* assign the application to a specific site in order to make it accessible. Assign the application to the required site           using the <a href=\"#operation/BindApplicationToSite\">Bind Application to Site API</a>.         </li>         <li>           In order to grant access to the application for specific entities (users/groups), you should assign the application to the access policy using the <a href=\"#tag/Access-and-Activity-Policies\">Access and Activity Policy API</a>         </li>       </ol>   </li> </ol>  ## Object Model The object model of the API is built around the following: <ol>   <li><a href=\"#tag/Sites\">Sites</a> - Site is a representation of the physical or virtual data center your applications reside in.</li>   <li><a href=\"#tag/Connectors\">Connectors</a> - A connector is a lightweight piece of software connecting your site to the Secure Access Cloud platform.</li>   <li><a href=\"#tag/Applications\">Applications</a>  - Application is the internal resource you would like to publish using Secure Access Cloud. </li>   <li>     <a href=\"#tag/Access-and-Activity-Policies\">Access and Activity Policies</a> - Secure Access Cloud continuously authorize each user request for the contextual access and activity,     in order to control access to resources and restrict user’s actions within resources, based on the user/device context (such as the user’s group membership, user’s location,     MFA status and managed/unmanaged device status) and the requested resource.   <li>     <a href=\"#tag/Cloud-Integration\">Cloud Integration</a> - Integration with Cloud Providers like Amazon Web Services to provide a smoother and cloud-native integration with SIEM solutions      and to allow access to resources based on their associated tags.   <li>     Logs - Secure Access Cloud internal logs for audit and forensics purposes:     <ol>       <li><a href=\"#tag/Audit-Logs\">Audit Logs</a> audit all operations done through the administration portal</li>       <li><a href=\"#tag/Web-Access-Logs\">Web Access Logs</a> audit any web access</li>       <li><a href=\"#tag/SSH-Logs\">SSH Logs</a> audit any SSH access</li>       <li><a href=\"#tag/RDP-Logs\">RDP Logs</a> audit any RDP access</li>       <li><a href=\"#tag/Forensics-Logs\">Forensics Logs</a> audit any user's access to any application as well as user's activity for SSH and RDP applications.</li>     </ol>   </li> </ol>   ## Authentication  Authentication is done using [OAuth2](https://tools.ietf.org/html/rfc6749) with the [Bearer authentication scheme](https://tools.ietf.org/html/rfc6750).  <!-- ReDoc-Inject: <security-definitions> -->  The Secure Access Cloud API is available to Secure Access Cloud users who have administrative privileges in their Secure Access Cloud tenant. An administrator should create an API client through the Secure Access Cloud Admin portal, check the ‘Allow access to Secure Access Cloud management API’ permission and copy the ‘Client Id’ and the ‘Client Secret’.  Retrieving the API access token is done using Basic-Authentication scheme, POST of a Base64 encoded Client-ID and Client-Secret: <B>   ``` curl -X POST \\  https://api.acme.luminatesec.com/v1/oauth/token \\  -u yourApiClientId:yourApiClientSecret   ``` </B>  This call returns the following JSON: {     \"access_token\":\"edfe22e3-eb4c-4c83-8ce3-3152e6a2XXX\",   \"expires_in\":86400,   \"scope\":\"luminate-scope\",   \"token_type\":\"Bearer\",   \"error\":\"\",   \"error_description\":\"\"}  All further API calls should include the ‘Authorization’ header with value “Bearer AccessToken”  For example: <B>   ```   curl -H \"Authorization: Bearer edfe22e3-eb4c-4c83-8ce3-3152e6a2XXX\" \"https://api.acme.luminatesec.com/v2/applications/d9f6ca17-9f2c-488c-aa86-51924a37092e\"   ``` </B>  ## Versioning and Compatibility  The latest Major Version is `v2`.  The Major Version is included in the URL path (e.g. /v2/applications ) and it denotes breaking changes to the API. Minor and Patch versions are transparent to the client.  ## Pagination   Some of our API responses are paginated, meaning that only a certain number of items are returned at a time.  The default number of items returned in a single page is 50.  You can override this by passing a size parameter to set the maximum number of results, but cannot exceed 100.  Specifying the page number sets the starting point for the result set, allowing you to fetch subsequent items  that are not in the initial set of results. The sort order for returned data can be controlled using the sort parameter.<br>  You can constrain the results by using a filter. <br><br>  **Note:** Most methods that support pagination use the approach specified above. However, some methods use varied   versions of pagination. The individual documentation for each API method is your source of truth for which pattern the method follows.  ## Auditing  All authentication operations and modify operations (POST, PUT, DELETE) are audited.   ## Rate-limiting The API has a rate limit of 5 requests per second. If you have hit the rate limit, then a ‘429’ status code will be returned. In such cases, you should back-off from submitting new requests for 1 second before resuming.  Note that rate-limitation applies to the accumulated requests of **all** of your clients. For example, if you have 6 clients submitting requests simultaneously at a rate of 1 request per second for each one then one of them is likely to get a 429 status code.  ## Support  For additional help you may refer to our support at https://support.luminate.io.  Each request submitted to the API returns a unique request ID that is generated by the API. The request ID will be returned in header `x-lum-request-id`. If you need to contact us about any specific request then this ID will serve as a reference to the given request. 
 *
 * API version: V2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type AuditLogSearchResults struct {
	// Total number of logs found that match the query.
	Hits int32 `json:"Hits,omitempty"`
	Logs []AuditLogResult `json:"Logs,omitempty"`
}

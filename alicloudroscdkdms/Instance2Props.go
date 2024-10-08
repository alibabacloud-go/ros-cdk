package alicloudroscdkdms


// Properties for defining a `Instance2`.
//
// See https://www.alibabacloud.com/help/ros/developer-reference/aliyun-dms-instance2
type Instance2Props struct {
	// Property databasePassword: The logon password of the database instance.
	DatabasePassword interface{} `field:"required" json:"databasePassword" yaml:"databasePassword"`
	// Property databaseUser: The logon username of the database instance.
	DatabaseUser interface{} `field:"required" json:"databaseUser" yaml:"databaseUser"`
	// Property dbaId: The user ID of the instance DBA role.
	//
	// You can call the ListUsers or GetInstance operation to obtain the value of this parameter.
	DbaId interface{} `field:"required" json:"dbaId" yaml:"dbaId"`
	// Property enableSellTrust: Whether the instance needs to enable secure hosting.
	//
	// - Y: Enable secure hosting mode.
	// - NULL or other: Do not enable secure hosting mode.
	EnableSellTrust interface{} `field:"required" json:"enableSellTrust" yaml:"enableSellTrust"`
	// Property envType: The type of the environment to which the database instance belongs.
	//
	// Valid values:
	// product: the production environment.
	// dev: the test environment.
	EnvType interface{} `field:"required" json:"envType" yaml:"envType"`
	// Property exportTimeout: The timeout period for exporting the database instance.
	//
	// Unit: seconds.
	ExportTimeout interface{} `field:"required" json:"exportTimeout" yaml:"exportTimeout"`
	// Property host: The endpoint of the database instance.
	Host interface{} `field:"required" json:"host" yaml:"host"`
	// Property instanceAlias: The alias of the database instance.
	//
	// The alias helps you quickly find the required
	// instance.
	InstanceAlias interface{} `field:"required" json:"instanceAlias" yaml:"instanceAlias"`
	// Property instanceSource: The source of the database instance.
	//
	// Valid values:
	// PUBLIC_OWN: an on-premises database built on the public network.
	// RDS: an ApsaraDB for RDS (RDS) instance.
	// ECS_OWN: an on-premises database built on an Elastic Compute Service (ECS) instance.
	// VPC_IDC: an on-premises database built in an Internet data center (IDC) in Virtual Private
	// Cloud (VPC).
	InstanceSource interface{} `field:"required" json:"instanceSource" yaml:"instanceSource"`
	// Property instanceType: The type of the database instance.
	//
	// Valid values: MySQL, SQLServer, PostgreSQL, Oracle, DRDS, OceanBase, Mongo, Redis.
	InstanceType interface{} `field:"required" json:"instanceType" yaml:"instanceType"`
	// Property networkType: The network type of the database instance.
	//
	// Valid values:
	// CLASSIC
	// VPC.
	NetworkType interface{} `field:"required" json:"networkType" yaml:"networkType"`
	// Property port: The connection port of the database instance.
	Port interface{} `field:"required" json:"port" yaml:"port"`
	// Property queryTimeout: The timeout period for querying the database instance.
	//
	// Unit: seconds.
	QueryTimeout interface{} `field:"required" json:"queryTimeout" yaml:"queryTimeout"`
	// Property safeRule: The security rule of the database instance.
	//
	// Enter the name of the security rule for
	// your enterprise.
	// Note To query a specified security rule, log on to the DMS Enterprise console and choose
	// System Management > Security Rules. The security rule appears in the security rule
	// list.
	SafeRule interface{} `field:"required" json:"safeRule" yaml:"safeRule"`
	// Property dataLinkName: The name of the data link for cross-database query.
	DataLinkName interface{} `field:"optional" json:"dataLinkName" yaml:"dataLinkName"`
	// Property ddlOnline: [Important] Specifies whether to enable the online data description language (DDL) service.
	//
	// Currently, this service is available only for the MySQL and PolarDB databases.
	// 0: The service is disabled.
	// 1: The native online DDL service prevails.
	// 2: Data change without table locking provided by DMS prevails.
	DdlOnline interface{} `field:"optional" json:"ddlOnline" yaml:"ddlOnline"`
	// Property ecsInstanceId: The ID of the ECS instance to which the database instance belongs.
	//
	// Note You must specify this parameter if the InstanceSource parameter is set to ECS_OWN.
	EcsInstanceId interface{} `field:"optional" json:"ecsInstanceId" yaml:"ecsInstanceId"`
	// Property ecsRegion: The region where the database instance resides.
	//
	// Note You must specify this parameter if the InstanceSource parameter is set to ECS_OWN or VPC_IDC.
	EcsRegion interface{} `field:"optional" json:"ecsRegion" yaml:"ecsRegion"`
	// Property enableSellCommon: Whether the instance needs to enable the security collaboration function.
	//
	// - Y: Enable.
	// - NULL or other: Not enable.
	EnableSellCommon interface{} `field:"optional" json:"enableSellCommon" yaml:"enableSellCommon"`
	// Property enableSellSitd: Y: enables the sensitive data protection feature NULL or other: disables the sensitive data protection feature.
	EnableSellSitd interface{} `field:"optional" json:"enableSellSitd" yaml:"enableSellSitd"`
	// Property enableSellStable: Whether the instance needs to enable the stable change function.
	//
	// - Y: Enable.
	// - NULL or other: Do not enable.
	EnableSellStable interface{} `field:"optional" json:"enableSellStable" yaml:"enableSellStable"`
	// Property sid: The system ID (SID) of the database instance.
	//
	// Note You must specify this parameter if the InstanceType parameter is set to PostgreSQL or Oracle.
	Sid interface{} `field:"optional" json:"sid" yaml:"sid"`
	// Property skipTest: Whether the instance ignores the test connectivity.
	//
	// The values are as follows:
	// - true: ignore
	// - false: do not ignore.
	SkipTest interface{} `field:"optional" json:"skipTest" yaml:"skipTest"`
	// Property templateId: The ID of the classification template.
	//
	// You can call the ListClassificationTemplates operation to query the template ID.
	TemplateId interface{} `field:"optional" json:"templateId" yaml:"templateId"`
	// Property templateType: The type of the classification template.
	//
	// You can call the ListClassificationTemplates operation to query the template type.
	TemplateType interface{} `field:"optional" json:"templateType" yaml:"templateType"`
	// Property tid: The ID of the tenant.
	//
	// Note To query the ID, log on to the DMS Enterprise console and choose System Management
	// > Instance Management or System Management > User Management. The ID of the tenant
	// appears in the Service Specification section.
	Tid interface{} `field:"optional" json:"tid" yaml:"tid"`
	// Property useDsql: Specifies whether to enable cross-database query for the database instance.
	//
	// Valid
	// values:
	// 0: disabled
	// 1: enabled.
	UseDsql interface{} `field:"optional" json:"useDsql" yaml:"useDsql"`
	// Property useSsl: Select whether to enable SSL connection.
	//
	// When enabled, DMS will connect to the database via SSL. The premise of using this parameter is that the database has enabled SSL connection. Currently, the following options are supported:
	// - 0: Default value (DMS and database service automatically negotiate whether to enable it) The system will automatically identify whether the target database has enabled SSL connection. If the database has enabled SSL connection, DMS will connect to the database via SSL; if the database has not enabled SSL connection, it will connect in a normal way.
	// - 1: Enable SSL connection. If the database has not enabled SSL connection, this parameter is invalid.
	// - -1: Disable SSL connection.
	// Note:
	// - This configuration item only appears when the database is MySQL or Redis.
	// - SSL (Secure Sockets Layer) can encrypt network connections at the transport layer to improve the security and integrity of communication data, but it will increase the network connection response time.
	UseSsl interface{} `field:"optional" json:"useSsl" yaml:"useSsl"`
	// Property vpcId: The ID of the VPC to which the database instance belongs.
	//
	// Note You must specify this parameter if the InstanceSource parameter is set to VPC_IDC.
	VpcId interface{} `field:"optional" json:"vpcId" yaml:"vpcId"`
}


/*
%%%%%%%%%%%%%插件版本：SG 1.0.0%%%%%%%%%%%%%
%%%%%%%%%%%%%该文件由工具自动生成，请不要手动修改%%%%%%%%%%%%%
%%%%%%%%%%%%%生成时间:2023-10-18 11:45:31%%%%%%%%%%%%%
*/

package gamedata

type ServerList struct {
	Id	int	"index"	//id
	Groupid	int	//所属组id
	Serverip	string	//服务器地址
	Paynoticecelp	string	//支付通知地址
	State	int	//服务器状态 1新服，2爆满
	Name	string	//服务器名称
	Time	int	//服务器开启时间
	Address	string	//地址
}

/*
%%%%%%%%%%%%%插件版本：SG 1.0.0%%%%%%%%%%%%%
%%%%%%%%%%%%%该文件由工具自动生成，请不要手动修改%%%%%%%%%%%%%
%%%%%%%%%%%%%生成时间:2024-03-21 16:47:31%%%%%%%%%%%%%
*/

package gamedata

type Shop struct {
	Id	int	"index"	//唯一id
	Name	string	//展示名称
	Costid	[]int	//消耗道具
	Costnum	[]int	//消耗数量
	Proptype	[]int	//道具类型
	Propnum	[]int	//道具数量
	Icon	string	//图标
	Limit	int	//限购
}

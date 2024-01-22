/*
%%%%%%%%%%%%%插件版本：SG 1.0.0%%%%%%%%%%%%%
%%%%%%%%%%%%%该文件由工具自动生成，请不要手动修改%%%%%%%%%%%%%
%%%%%%%%%%%%%生成时间:2023-12-14 09:51:41%%%%%%%%%%%%%
*/

package gamedata

type Chapter_stars struct {
	Id	int	"index"	//关卡id
	Star	[][]int	//星级对应时间范围
	Pos	[]int	//星级对应金币奖励
	Instruction	string	//说明
}

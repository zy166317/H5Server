/*
%%%%%%%%%%%%%插件版本：SG 1.0.0%%%%%%%%%%%%%
%%%%%%%%%%%%%该文件由工具自动生成，请不要手动修改%%%%%%%%%%%%%
%%%%%%%%%%%%%生成时间:2024-03-21 16:47:31%%%%%%%%%%%%%
*/

package gamedata

type Chapter_stars struct {
	Id	int	"index"	//关卡id
	Star	[][]int	//星级对应时间范围
	Rewards	[][]int	//星级对应金币奖励
	Alltime	int	//总时间
	Instruction	string	//说明
}

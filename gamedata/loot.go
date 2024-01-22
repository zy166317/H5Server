/*
%%%%%%%%%%%%%插件版本：SG 1.0.0%%%%%%%%%%%%%
%%%%%%%%%%%%%该文件由工具自动生成，请不要手动修改%%%%%%%%%%%%%
%%%%%%%%%%%%%生成时间:2023-12-14 09:51:41%%%%%%%%%%%%%
*/

package gamedata

type Loot struct {
	Id	int	"index"	//关卡id
	Propaweight	int	//道具a掉落概率
	Dropacondition	int	//道具a掉落条件
	Asubweight	int	//道具a概率每次递减
	Aaddcondition	int	//道具a掉落条件递增
	Propbweight	int	//道具b掉落概率
	Dropbcondition	int	//道具b掉落条件
	Bsubweight	int	//道具b概率每次递减
	Baddcondition	int	//道具b掉落条件递增
	Propcweight	int	//道具c掉落条件
	Dropccondition	int	//道具c掉落条件
	Csubweight	int	//道具c概率每次递减
	Caddcondition	int	//道具c掉落条件递增
	Instruction	string	//说明
}

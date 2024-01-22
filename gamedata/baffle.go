/*
%%%%%%%%%%%%%插件版本：SG 1.0.0%%%%%%%%%%%%%
%%%%%%%%%%%%%该文件由工具自动生成，请不要手动修改%%%%%%%%%%%%%
%%%%%%%%%%%%%生成时间:2023-12-14 09:51:41%%%%%%%%%%%%%
*/

package gamedata

type Baffle struct {
	Id	int	"index"	//唯一id
	Functionname	string	//道具名称
	Types	int	//对应方块类型
	Probability	int	//触发概率
	Subconditiona	int	//降低道具a掉落条件
	Subconditionb	int	//降低道具b掉落条件
	Subconditionc	int	//降低道具c掉落条件
}

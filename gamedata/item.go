/*
%%%%%%%%%%%%%插件版本：SG 1.0.0%%%%%%%%%%%%%
%%%%%%%%%%%%%该文件由工具自动生成，请不要手动修改%%%%%%%%%%%%%
%%%%%%%%%%%%%生成时间:2024-03-21 16:47:31%%%%%%%%%%%%%
*/

package gamedata

type Item struct {
	Id	int	"index"	//唯一id
	Functionname	string	//道具名称
	Price	[]int	//消耗道具
	Type	int	//道具类型
	Args1	[][]int	//参数一
	Args2	float32	//参数二
}

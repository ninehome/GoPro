
<template>

    <el-menu background-color="#545c64" text-color="#fff" active-text-color="#ffd04b" default-active="1-4-1"
        class="el-menu-vertical-demo" @open="handleOpen" @close="handleClose" :collapse="isCollapse">
        <h1 class="LeftTitle"> {{ isCollapse?'后台':'后台管理系统' }}</h1>
        <!-- :index 等于 v-bingde-->
        <el-menu-item  @click="homeClick(item)" v-for="item in noChildren" :key="item.name" :index="item.name">
            <!-- i图标 -->
            <i :class="'el-icon-${item.icon}'"></i>

            <span slot="title">{{ item.lable }}</span>
        </el-menu-item>

        




        <el-submenu v-for="sonitem in hasChildren" :key="sonitem.lable" :index="sonitem.lable">

            <template slot="title">
                <i class="el-icon-location"></i>
                <span slot="title">{{ sonitem.lable }}</span>
            </template>



            <!-- :index 属性用来保证唯一性 -->
            <el-menu-item-group  v-for="subItem in sonitem.children" :key="subItem.name" :index="subItem.path">

                <el-menu-item  @click="subClick(subItem)"  :index="subItem.path">{{ subItem.name}}</el-menu-item>

            </el-menu-item-group>

        </el-submenu>




    </el-menu>


</template>

 <!-- 使用less css 插件，scoped-->只作用于当前 -->
<style lang="less" scoped>
.el-menu-vertical-demo:not(.el-menu--collapse) {
    width: 200px;
    background-color: rgb(18, 21, 22);
    text-emphasis-color: rgb(227, 235, 240);
    height: 100vh;
    padding: 0;
    border-right: none;

    .LeftTitle{
        color: #fff;
        text-align: center;
        line-height: 48px;
        font-size: 16px;
        font-weight: 400;
    }

}


</style>
  
<script>
export default {
    data() {
        return {
            // isCollapse: false, 底部已经定义，这里注释，避免重复定义
            menuData: [
                {
                    path: '/',
                    name: 'home',
                    lable: '首页',
                    icon: 's-home',
                    url: 'Home/Home'
                },
                {
                    path: '/mall',
                    name: 'mall',
                    lable: '商品管理',
                    icon: 'video-play',
                    url: 'MallManage/MallManage'
                },
                {
                    path: '/user',
                    name: 'user',
                    lable: '用户管理',
                    icon: 'user',
                    url: 'UserManage/UserManage'
                },
                {
                    lable: '其他',
                    icon: 'location',
                    children: [
                        {
                            path: '/page1',
                            name: 'page1',
                            lable: '页面1',
                            icon: 'stting',
                            url: 'Other/PageOne'
                        },
                        {
                            path: '/page2',
                            name: 'page2',
                            lable: '页面2',
                            icon: 'stting',
                            url: 'Other/PageTwo'
                        }
                    ]

                }
            ]
        };
    },
    methods: {
        handleOpen(key, keyPath) {
            console.log(key, keyPath);
        },
        handleClose(key, keyPath) {
            console.log(key, keyPath);
        },
        homeClick(item){
              //当前页面 地址 和 点击传入地址不一致才需要jump
            if(this.$route !== item.path){
                this.$router.push(item.path)
            }
        

        },
        subClick(subItem){
            // console.log(subItem)
            if(this.$route !== subItem.path){
                this.$router.push(subItem.path)
            }
            
        }
    },
    // 数据分组过滤
    computed: {
        //   没有子菜单
        noChildren() {
            //遍历数据返回有子菜单的 返回集合数据
            return this.menuData.filter(
                item => !item.children
            )
        },

        //   有子菜单
        hasChildren() {
            //遍历数据返回有子菜单的
            return this.menuData.filter(
                item => item.children
            )
        },
        isCollapse(){
            return this.$store.state.tab.isCollapse

        }
    }
}
</script>


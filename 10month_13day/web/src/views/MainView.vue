<template>
    <div>

        <el-container>

            <!-- auto 宽度自适应 -->
            <el-aside width="auto">

                <!--  用中划线 小驼峰 来替代ComonAsid ,也可直接使用ComonAsid-->
                <comon-asid> </comon-asid>
            </el-aside>
            <el-container>
                <el-header>
                    <comonHead></comonHead>

                </el-header>


                <el-main>

        
                    <!-- 路由匹配到的组件将渲染到这里 -->
                    <router-view></router-view>

                </el-main>
            </el-container>
        </el-container>
    </div>
</template>


<script>

import ComonAsid from '../components/CommonAsid.vue'
import comonHead from '../components/CommonHead.vue'
export default {
    data() {
        return {}
    },

    // 引入左边选择组建
    components: {
        ComonAsid,
        comonHead
    },

}
</script>

<style scoped>
.el-header {

    padding: 0;

}
</style>


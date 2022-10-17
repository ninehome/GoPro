<template>
  <div>
    <el-row>

      <el-col :span="8">

        <!-- 左边个人信息 -->
        <el-card class="box-card">
          <div class="user">
            <img class="img" src="../assets/images/ic_001.jpg" />
            <div class="userinfo">
              <p class="admin">Admin</p>
              <p class="accser">超级管理员</p>
            </div>
          </div>


          <div class="login-info">
            <p class="logintime">上次登录时间:<span>2022-10-17</span></p>
            <p class="loginaddr">上次登录地点:<span>马尼拉</span></p>
          </div>

        </el-card>



        <!-- 左边表格数据 -->
        <el-card class="datamenu">

          <el-table :data="tableData" stripe style="width: 100%">

            <el-table-column v-for="(val,key) in tabLabel " :prop="key" :label="val" />


          </el-table>



        </el-card>




      </el-col>






      <!-- 右边UI -->
      <el-col :span="16">
        <div class="num">
          <el-card class="rihtCard" v-for="c in countData" :key="c.name" :body-style="{display : 'flex',padding:0}">
            <i class="icon" :class="c.icon" :style="{ background : c.color}"></i>
            <div class="detali">
              <p>{{c.value}}</p>
              <p>{{c.name}}</p>
            </div>

          </el-card>

        </div>
      </el-col>


    </el-row>
  </div>




</template>












<style lang="less" scoped>
.user {
  border-bottom: 1px solid #999999;
  display: flex;
  align-items: center; //垂直居中
  padding-bottom: 10px;
  padding-top: 10px;
  padding-left: 10px;

  .img {

    width: 120px;
    height: 120px;
    border-radius: 20%;
  }

  .userinfo {
    line-height: 0.5;
    margin-left: 20px;


    .admin {
      font-size: 20px;
    }

    .accser {
      font-size: 12px;
    }


  }
}

.login-info {
  line-height: 0.5;
  margin-top: 20px;

  .logintime {
    margin-top: 5px;
    font-size: 15px;

  }

  .loginaddr {
    color: #999999;
  }

}

.num {

  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;



  .icon {

    width: 80px;
    height: 80px;
    text-align: center;
    line-height: 80px;
    background-color: #fff;
    border-radius: 5px;
  }

  .detali {
    display: flex;
    flex-direction: column;
    justify-content: center;
    padding-left: 10px;

  }

  .rihtCard {
    width: 30%;

  }

}
</style>








<script>


import { getData } from '../api'


export default {
  data() {
    return {
      tableData: [{
        date: '2016-05-02',
        name: '王小虎',
        address: '上海市普陀区'
      },
      {
        date: '2016-05-04',
        name: '王小虎',
        address: '南京'
      },
      {
        date: '2016-05-01',
        name: '王小虎',
        address: '长沙'
      },
      {
        date: '2016-05-03',
        name: '王小虎',
        address: '武汉'
      }],

      tabLabel: {
        date: '时间',
        name: '姓名',
        address: '地址',

      },

      countData: [{
        value: '13131',
        name: '今日收藏',
        color: '#5ab1ef',
        icon: "el-icon-cherry"

      },
      {
        value: '13131',
        name: '今日支付67',
        color: '#5ab1ef',
        icon: "el-icon-dish"
      },
      {
        value: '13131',
        name: '本月支付',
        color: '#5ab1ef',
        icon: "el-icon-check"

      },
      {
        value: '13131',
        name: '今日支付2',
        color: '#5ab1ef',
        icon: "el-icon-dish"
      },
      {
        value: '13131',
        name: '本月支付4',
        color: '#5ab1ef',
        icon: "el-icon-check"

      },
      {
        value: '13131',
        name: '今日支付5',
        color: '#5ab1ef',
        icon: "el-icon-dish"
      }


      ]

    }
  },
  mounted(){
    getData().then((data)=>{
      console.log(data)
    })

  }
}
</script>
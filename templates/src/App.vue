<template>
  <div id="app">
    <el-container>
      <el-header>Header</el-header>
      <el-container>
        <el-aside width="200px">
          <el-menu :default-active="defaultActive" router background-color="#e9eef3">
            <sider-menu :menuList="menuList"></sider-menu>
          </el-menu>
        </el-aside>
        <el-main>
          <router-view/>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script>
import SiderMenu from "components/siderMenu";
import { menuList } from "./const";

export default {
  name: "App",
  data() {
    return {
      menuList: []
    };
  },
  computed: {
    defaultActive() {
      return this.$route.path;
    }
  },
  components: {
    SiderMenu
  },
  mounted() {
    this.menuList = menuList;
  }
};
</script>


<style>
#app {
  height: 100%;
  overflow: hidden;
}
.el-container.is-vertical {
  height: 100%;
}
.el-header {
  background-color: #b3c0d1;
  color: #333;
  text-align: center;
  line-height: 60px;
}

.el-aside {
  background-color: #e9eef3;
  color: #333;
  line-height: 200px;
  height: 100%;
  overflow: auto;
}

.el-main {
  color: #333;
}

body > .el-container {
  margin-bottom: 40px;
}
</style>

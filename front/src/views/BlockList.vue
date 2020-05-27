<template>
  <el-main>
    <div style="display: flex;justify-content: center;">
      <!-- <el-main> -->
      <div style="width:900px">
        <blockhead
          v-for="block in blocks"
          :key="block"
          :block="block"
          style="margin-bottom:10px"
          v-on:dblclick.native="detail(block)"
        ></blockhead>
      </div>
    </div>
  </el-main>
</template>

<script>
import blockhead from "../components/BlockHead.vue";
export default {
  components: {
    blockhead
  },
  props: {
    blocks: Array
  },
  data() {
    return {};
  },
  methods: {
    init() {
      if (localStorage.ledger) {
        this.blocks = JSON.parse(localStorage.ledger);
        localStorage.removeItem("block");
      } else {
        this.$message({
          message: "没有发现区块数据，请重新上传账本",
          type: "warnning"
        });
      }
    },
    detail(block) {
      localStorage.block = JSON.stringify(block);
      if (block.type == 1) {
        this.$router.push({
          name: "Config"
        });
      } else if (block.type == 0) {
        this.$router.push({
          name: "Transaction"
          // params: {
          //   block: block
          // }
        });
      } else {
        this.$message({
          message: "不支持的区块类型",
          type: "warnning"
        });
      }
    }
  },
  mounted() {
    this.init();
  }
};
</script>

<style>
</style>
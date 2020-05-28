<template>
  <el-main>
    <div style="display: flex;justify-content: center;">
      <!-- <el-main> -->
      <div style="width:900px">
        <blockhead
          :v-if="blocks"
          v-for="block in blocks"
          :key="block"
          :block="block"
          style="margin-bottom:10px"
          v-on:dblclick.native="detail(block)"
        ></blockhead>
        <el-pagination
          :v-if="pagination"
          @current-change="handleCurrentChange"
          :current-page.sync="curpage"
          :page-size="20"
          layout="total, prev, pager, next"
          :total="total"
        >
        </el-pagination>
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
    return {
      curpage: 0,
      total: 0,
      pagination: false,
      key: ""
    };
  },
  methods: {
    init() {
      if (localStorage.ledger) {
        const ledger = JSON.parse(localStorage.ledger);
        this.blocks = ledger.blocks;
        this.pagination = ledger.pagination;
        this.curpage = ledger.cur_page;
        this.total = ledger.total;
        this.key = ledger.key;
        this.localStorage.removeItem("block");
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
    },
    handleCurrentChange(page) {
      const action = `${this.$consts.ServerAddrPre}/ledger/blocks`;
      this.$axios
        .get(action, {
          params: {
            key: this.key,
            page: page
          }
        })
        .then(rsp => {
          console.log(rsp);
          if (rsp.data.code === 200) {
            this.blocks = rsp.data.data.blocks;
            localStorage.ledger = JSON.stringify(rsp.data.data);
          } else {
            this.$message({
              message: "请求数据出错" + rsp.data.msg,
              type: "warnning"
            });
          }
        })
        .catch(err => {
          this.$message.error(err.message);
        });
    }
  },
  mounted() {
    this.init();
  }
};
</script>

<style>
</style>
<template>
  <el-main>
    <el-upload
      class=""
      drag
      :action="action"
      name="ledger"
      :on-success="success"
      :on-error="onError"
    >
      <i class="el-icon-upload"></i>
      <div class="el-upload__text">将账本文件拖到此处，或<em>点击上传</em></div>
      <div class="el-upload__tip" slot="tip">
        上传 fabric peer 节点存储的账本文件（eg:blockfile_000000）
      </div>
    </el-upload>
  </el-main>
</template>

<script>
export default {
  data() {
    return {
      action: `${this.$consts.ServerAddrPre}/ledger/file`
    };
  },

  methods: {
    success(response, file, fileList) {
      // console.log(response);
      localStorage.removeItem("ledger");
      if (response.code == 200) {
        localStorage.ledger = JSON.stringify(response.data);
        this.$router.push({
          name: "BlockList"
        });
      } else {
        this.$message({
          message:
            "账本解析错误，请检查上传文件是否正确（" + response.msg + ")",
          type: "warnning"
        });
      }
    },
    onError(err, file, fileList) {
      this.$message.error("上传失败");
      localStorage.ledger = null;
    }
  }
};
</script>

<style>
</style>
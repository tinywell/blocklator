<template>
  <el-main>
    <el-upload
      class=""
      drag
      :action="action"
      name="block"
      :on-success="success"
      :on-error="onError"
    >
      <i class="el-icon-upload"></i>
      <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
      <div class="el-upload__tip" slot="tip">
        上传单个区块文件，通过 'peer channel fetch'
        命令或其他方式从网络中拉取的区块文件
      </div>
    </el-upload>
  </el-main>
</template>

<script>
export default {
  data() {
    return {
      action: `${this.$consts.ServerAddrPre}/block/file`
    };
  },
  methods: {
    success(response, file, fileList) {
      console.log(response);
      if (response.code == 200) {
        if (response.data.type == 1) {
          this.$router.push({
            name: "Config",
            params: {
              block: response.data
            }
          });
        } else if (response.data.type == 0) {
          this.$router.push({
            name: "Transaction",
            params: {
              block: response.data
            }
          });
        }
      } else {
        this.$message({
          message: "区块解析错误，请检查上传文件是否正确",
          type: "warnning"
        });
      }
    },
    onError(err, file, fileList) {
      // console.log(this.action);
      // console.log(err);
      this.$message.error("上传失败");
    }
  }
};
</script>

<style>
</style>
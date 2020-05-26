<template>
  <div>
    <el-row>
      <el-col :span="2" class="cert_label">CN：</el-col>
      <el-col :span="20" class="code_text">{{ cert.cn }}</el-col>
    </el-row>
    <el-row>
      <el-col :span="2" class="cert_label">O：</el-col>
      <el-col :span="20" class="code_text">{{ cert.org }}</el-col>
    </el-row>
    <el-row v-if="cert.ou != ''">
      <el-col :span="2" class="cert_label">OU：</el-col>
      <el-col :span="20" class="code_text">{{ cert.ou }}</el-col>
    </el-row>
    <el-row>
      <el-col class="code_text">{{ cert.pem }}</el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  props: {
    cert: Object
  }
};
</script>

<style lang="less">
.cert_label {
  text-align: right;
}

.code_text {
  font-family: Menlo, Monaco, Consolas, Courier, monospace;
  font-size: smaller;
  white-space: pre-wrap;
  word-break: break-all;
}
</style>
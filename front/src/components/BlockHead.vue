<template>
  <el-card class="box-card">
    <el-row class="row">
      <el-col :span="4" class="label">区块号：</el-col>
      <el-col :span="20" class="text">{{ block.block_num }}</el-col>
    </el-row>
    <el-row>
      <el-col :span="4" class="label">通道：</el-col>
      <el-col :span="20" class="text">{{ block.channel }}</el-col>
    </el-row>
    <el-row>
      <el-col :span="4" class="label">区块哈希：</el-col>
      <el-col :span="20" class="text code_text">{{ block.hash }}</el-col>
    </el-row>
    <el-row>
      <el-col :span="4" class="label">前区块哈希：</el-col>
      <el-col :span="20" class="text code_text">{{ block.pre_hash }}</el-col>
    </el-row>
    <el-row class="row">
      <el-col :span="4" class="label">区块类型：</el-col>
      <el-col :span="20" class="text">
        {{ this.$consts.BlockTypeOptsMap[block.type] }}
      </el-col>
    </el-row>
  </el-card>
</template>

<script>
export default {
  props: {
    block: Object
  }
};
</script>

<style lang="less">
.row {
  display: flex;
  justify-items: center;
  align-content: center;
}

.label {
  // width: 120px;
  text-align: right;
}

.text {
  text-align: left;
}

.code_text {
  white-space: pre-wrap;
  font-family: Menlo, Monaco, Consolas, Courier, monospace;
  font-size: medium;
}
</style>
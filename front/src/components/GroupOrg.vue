<template>
  <div class="org-card">
    <div class="card_name">{{ org.name }}</div>
    <div class="card_split"></div>
    <el-tabs tab-position="left" style="height: 300px;">
      <el-tab-pane label="根证书">
        <div class="card_certs">{{ org.root_cert }}</div>
      </el-tab-pane>
      <el-tab-pane label="TLS根证书">
        <div class="card_certs">{{ org.tls_root_cert }}</div>
      </el-tab-pane>
      <el-tab-pane label="管理员证书">
        <div class="card_certs">{{ org.admin }}</div>
      </el-tab-pane>
    </el-tabs>
    <div class="card_endpoints">
      Endpoints:
      <el-tag
        v-for="endpoint in org.endpoints"
        :key="endpoint"
        type="success"
        size="small"
        class="card_endpoint"
      >
        {{ endpoint }}
      </el-tag>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    org: Object
  }
};
</script>

<style lang="less">
.org-card {
  background-color: #fff;
  border-radius: 15px;
  box-shadow: 0 0 5px rgba(fade(#000, 1%));
  display: flex;
  padding: 20px 20px;
  flex-direction: column;
}

.card_name {
  font-size: 18px;
  text-align: center;
  font-weight: bold;
  font-family: Menlo, Monaco, Consolas, Courier, monospace;
}

.card_certs {
  padding-top: 10px;
  text-align: left;
  white-space: pre-wrap;
  font-family: Menlo, Monaco, Consolas, Courier, monospace;
  font-size: small;
}

.card_endpoints {
  text-align: left;
  margin: 20px 20px;
  .card_endpoint {
    margin-right: 20px;
    margin-left: 10px;
    font-family: Menlo, Monaco, Consolas, Courier, monospace;
  }
}

.card_split {
  margin: 10px 0;
  width: 100%;
  border-bottom: 1px solid #e9edf0;
}
</style>
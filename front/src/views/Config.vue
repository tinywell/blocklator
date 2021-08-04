<template>
  <div style="display: flex; justify-content: center">
    <!-- <el-main> -->
    <div style="width: 900px">
      <blockhead :block="block"></blockhead>
      <el-card
        class="box-card"
        v-if="Object.keys(block.config.consortium_orgs).length > 0"
      >
        <!-- <div slot="header" class="clearfix">
          <span>联盟组织</span>
        </div> -->
        <el-collapse>
          <el-collapse-item
            v-for="(orgs, name) in block.config.consortium_orgs"
            :key="(orgs, name)"
          >
            <template slot="title">
              <div class="row">
                <div>联盟：</div>
                <div>{{ name }}</div>
              </div>
            </template>
            <ul class="org_list">
              <li class="org" v-for="org in orgs" :key="org">
                <el-popover
                  trigger="hover"
                  placement="bottom"
                  width="700"
                  show="show"
                >
                  <grouporg :org="org" org_type="orderer"></grouporg>
                  <el-card shadow="hover" slot="reference">
                    {{ org.name }}
                  </el-card>
                </el-popover>
              </li>
            </ul>
            <!-- <grouporg :org="org" org_type="peer"></grouporg> -->
          </el-collapse-item>
        </el-collapse>
      </el-card>
      <el-card class="box-card" v-if="block.config.application_orgs.length > 0">
        <div slot="header" class="clearfix">
          <span>应用组织</span>
        </div>
        <ul class="org_list">
          <li
            class="org"
            v-for="org in block.config.application_orgs"
            :key="org"
          >
            <el-popover
              trigger="hover"
              placement="bottom"
              width="700"
              show="show"
            >
              <grouporg :org="org" org_type="peer"></grouporg>
              <!-- <el-tag type="success" slot="reference">{{ org.name }}</el-tag> -->
              <el-card shadow="hover" slot="reference">
                {{ org.name }}
              </el-card>
            </el-popover>
          </li>
        </ul>
      </el-card>
      <el-card class="box-card">
        <div slot="header" class="clearfix">
          <span>共识组织</span>
        </div>
        <ul class="org_list">
          <li class="org" v-for="org in block.config.orderer_orgs" :key="org">
            <el-popover
              trigger="hover"
              placement="bottom"
              width="700"
              show="show"
            >
              <grouporg :org="org" org_type="orderer"></grouporg>
              <el-card shadow="hover" slot="reference">
                {{ org.name }}
              </el-card>
            </el-popover>
          </li>
        </ul>
      </el-card>
      <el-row class="box-card">
        <el-col :span="12">
          <el-card class="param-card1" style="margin-right: 5px; height: 350px">
            <div slot="header" class="clearfix">
              <span>共识参数</span>
            </div>
            <el-row class="row">
              <el-col :span="12" class="label">共识类型：</el-col>
              <el-col :span="12" class="text">{{
                block.config.consensus.type
              }}</el-col>
            </el-row>
            <el-row class="row">
              <el-col :span="12" class="label">区块最大交易数：</el-col>
              <el-col :span="12" class="text">{{
                block.config.consensus.max_message_count
              }}</el-col>
            </el-row>
            <el-row class="row">
              <el-col :span="12" class="label"
                >区块最大绝对值 absolute：</el-col
              >
              <el-col :span="12" class="text"
                >{{ block.config.consensus.absolute_max_bytes }} (Byte)</el-col
              >
            </el-row>
            <el-row class="row">
              <el-col :span="12" class="label"
                >区块最大偏好值 preferred：</el-col
              >
              <el-col :span="12" class="text"
                >{{ block.config.consensus.preferred_max_bytes }} (Byte)</el-col
              >
            </el-row>
            <el-row class="row">
              <el-col :span="12" class="label">出块超时时间：</el-col>
              <el-col :span="12" class="text"
                >{{ block.config.consensus.batch_time_out }}
              </el-col>
            </el-row>
            <el-row v-if="block.config.consensus.type == 'etcdraft'">
              <el-col :span="12" class="label">raft consenters：</el-col>
              <el-col :span="12" class="text">
                <el-tag
                  v-for="consenter in block.config.consensus.raft_metadata
                    .consenters"
                  :key="consenter"
                  size="small"
                  style="margin: 2px 0px; width: 200px; text-align: center"
                >
                  {{ consenter.host }}:{{ consenter.port }}
                </el-tag>
              </el-col>
            </el-row>
          </el-card>
        </el-col>
        <el-col :span="12">
          <el-card style="margin-left: 5px; height: 350px">
            <div slot="header" class="clearfix">
              <span>网络参数</span>
            </div>
            <el-row class="row">
              <el-col :span="12" class="label">hash 算法：</el-col>
              <el-col :span="12" class="text"
                >{{ block.config.values.hashing_algorithm }}
              </el-col>
            </el-row>
            <el-row class="row">
              <el-col :span="12" class="label">联盟：</el-col>
              <el-col :span="12" class="text"
                >{{ block.config.values.consortium }}
              </el-col>
            </el-row>
            <el-row class="row" style="margin-bottom: 5px">
              <el-col :span="12" class="label">共识服务节点地址：</el-col>
              <el-col
                :span="12"
                class="text"
                style="overflow-x: hidden; overflow-y: zuto; height: 100px"
              >
                <el-tag
                  v-for="addr in block.config.values.orderer_addresses"
                  :key="addr"
                  size="small"
                  style="margin: 2px 0px; width: 200px; text-align: center"
                >
                  {{ addr }}
                </el-tag>
              </el-col>
            </el-row>
            <el-row class="row">
              <el-col :span="12" class="label">Channel Capabilities：</el-col>
              <el-col :span="12" class="text">
                <el-tag
                  v-for="cap in block.config.consensus.capabilities"
                  :key="cap"
                  size="small"
                  style="margin: 2px 0px; width: 60px; text-align: center"
                >
                  {{ cap }}
                </el-tag>
              </el-col>
            </el-row>
            <el-row class="row">
              <el-col :span="12" class="label">Orderer Capabilities：</el-col>
              <el-col :span="12" class="text">
                <el-tag
                  v-for="cap in block.config.values.capabilities"
                  :key="cap"
                  size="small"
                  style="margin: 2px 0px; width: 60px; text-align: center"
                >
                  {{ cap }}
                </el-tag>
              </el-col>
            </el-row>
            <el-row class="row">
              <el-col :span="12" class="label"
                >Applicaiton Capabilities：</el-col
              >
              <el-col
                :span="12"
                class="text"
                v-if="block.config.application_values"
              >
                <el-tag
                  v-for="cap in block.config.application_values.capabilities"
                  :key="cap"
                  size="small"
                  style="margin: 2px 0px; width: 60px; text-align: center"
                >
                  {{ cap }}
                </el-tag>
              </el-col>
            </el-row>
          </el-card>
        </el-col>
      </el-row>
    </div>
    <!-- </el-main> -->
  </div>
</template>

<script>
import grouporg from "../components/GroupOrg.vue";
import blockhead from "../components/BlockHead.vue";
export default {
  components: { grouporg, blockhead },
  data() {
    return {
      block: {
        // block_num: 0,
        // hash:
        //   "459D18A9DE01A0577EEF4B04F14902E714DED91B80193D33AB101D07054D7FD9",
        // pre_hash: "",
        // channel: "mychannel",
        // type: 1,
        // config: {
        //   orderer_orgs: [
        //     {
        //       name: "OrdererMSP",
        //       root_cert:
        //         "-----BEGIN CERTIFICATE-----\nMIICPDCCAeOgAwIBAgIQZSa5bKtAfeX9W+5olwNjSjAKBggqhkjOPQQDAjBpMQsw\nCQYDVQQGEwJVUzETMBEGA1UECBMKQ2FsaWZvcm5pYTEWMBQGA1UEBxMNU2FuIEZy\nYW5jaXNjbzEUMBIGA1UEChMLZXhhbXBsZS5jb20xFzAVBgNVBAMTDmNhLmV4YW1w\nbGUuY29tMB4XDTE5MDQxNjA2MTEwMFoXDTI5MDQxMzA2MTEwMFowaTELMAkGA1UE\nBhMCVVMxEzARBgNVBAgTCkNhbGlmb3JuaWExFjAUBgNVBAcTDVNhbiBGcmFuY2lz\nY28xFDASBgNVBAoTC2V4YW1wbGUuY29tMRcwFQYDVQQDEw5jYS5leGFtcGxlLmNv\nbTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABIv6ieqkILRw83QjpURhZBidgfSN\nmI1D9aIMxuvvwoQdngRAYFViLRVo/z7DCf29DWGnhg/GquGpizHgzHWjXIyjbTBr\nMA4GA1UdDwEB/wQEAwIBpjAdBgNVHSUEFjAUBggrBgEFBQcDAgYIKwYBBQUHAwEw\nDwYDVR0TAQH/BAUwAwEB/zApBgNVHQ4EIgQg79vAPjSJSR9HLdyKuiNWTQ3lT1wB\nUl3V+cOTVCCnIu4wCgYIKoZIzj0EAwIDRwAwRAIgWY5s6KP+mDPQbA0ROr2NcucC\nVRHb+IMttx2IIMaSJ/8CIA25kF7Dapj3WU0euF2AOVzQLyigeMe9Ck/HiTiWJbsy\n-----END CERTIFICATE-----\n",
        //       tls_root_cert:
        //         "-----BEGIN CERTIFICATE-----\nMIICQzCCAeqgAwIBAgIRANHC4rQsUx9/OD2sjAo5KCgwCgYIKoZIzj0EAwIwbDEL\nMAkGA1UEBhMCVVMxEzARBgNVBAgTCkNhbGlmb3JuaWExFjAUBgNVBAcTDVNhbiBG\ncmFuY2lzY28xFDASBgNVBAoTC2V4YW1wbGUuY29tMRowGAYDVQQDExF0bHNjYS5l\neGFtcGxlLmNvbTAeFw0xOTA0MTYwNjExMDBaFw0yOTA0MTMwNjExMDBaMGwxCzAJ\nBgNVBAYTAlVTMRMwEQYDVQQIEwpDYWxpZm9ybmlhMRYwFAYDVQQHEw1TYW4gRnJh\nbmNpc2NvMRQwEgYDVQQKEwtleGFtcGxlLmNvbTEaMBgGA1UEAxMRdGxzY2EuZXhh\nbXBsZS5jb20wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATh/V86Uclf+pKrLtl1\nrIEx9lSn/+v0fBRg3FLl6BWeCecw9c2ZccA0sSU6lg/jEMotOzEI2KaofQH40iIO\ngI73o20wazAOBgNVHQ8BAf8EBAMCAaYwHQYDVR0lBBYwFAYIKwYBBQUHAwIGCCsG\nAQUFBwMBMA8GA1UdEwEB/wQFMAMBAf8wKQYDVR0OBCIEIMZqNe7wu+eRVSqyFbRL\nAugjvnNIbTSM0httbsR6bhuMMAoGCCqGSM49BAMCA0cAMEQCIA/IZcnf3lFONTnw\nXbAxkGpB47HrF1pEHVoelHowK5AFAiAubVYjW+WLYbBEw7NJBJbx0DZaZpB+Y1Ko\n92zJjm2bzQ==\n-----END CERTIFICATE-----\n",
        //       admin:
        //         "-----BEGIN CERTIFICATE-----\nMIICCjCCAbGgAwIBAgIRAPs5wlpYksqMUsBMIDkjKOIwCgYIKoZIzj0EAwIwaTEL\nMAkGA1UEBhMCVVMxEzARBgNVBAgTCkNhbGlmb3JuaWExFjAUBgNVBAcTDVNhbiBG\ncmFuY2lzY28xFDASBgNVBAoTC2V4YW1wbGUuY29tMRcwFQYDVQQDEw5jYS5leGFt\ncGxlLmNvbTAeFw0xOTA0MTYwNjExMDBaFw0yOTA0MTMwNjExMDBaMFYxCzAJBgNV\nBAYTAlVTMRMwEQYDVQQIEwpDYWxpZm9ybmlhMRYwFAYDVQQHEw1TYW4gRnJhbmNp\nc2NvMRowGAYDVQQDDBFBZG1pbkBleGFtcGxlLmNvbTBZMBMGByqGSM49AgEGCCqG\nSM49AwEHA0IABCrSwC/TaNTZW7sVHx51j3E+VnSFbjWu/Kg4U5d3eiqzCUIQ8y09\nH9csQ59kTCkkLcUn/9cGx1Y+N+Q52bkukRmjTTBLMA4GA1UdDwEB/wQEAwIHgDAM\nBgNVHRMBAf8EAjAAMCsGA1UdIwQkMCKAIO/bwD40iUkfRy3cirojVk0N5U9cAVJd\n1fnDk1QgpyLuMAoGCCqGSM49BAMCA0cAMEQCICaOd8SzAZXeDi65LGW9cOrEjfI/\nAGvtD+QKU3TJvwHcAiAE53TDKOnDQQZi/CHhOVuOjON1SY6n1sGyEyJeO11z6g==\n-----END CERTIFICATE-----\n"
        //     }
        //   ],
        //   application_orgs: [
        //     {
        //       name: "Org1MSP",
        //       root_cert:
        //         "-----BEGIN CERTIFICATE-----\nMIICUTCCAfegAwIBAgIQVKbHrQz4eCVmoqjDSMKN+TAKBggqhkjOPQQDAjBzMQsw\nCQYDVQQGEwJVUzETMBEGA1UECBMKQ2FsaWZvcm5pYTEWMBQGA1UEBxMNU2FuIEZy\nYW5jaXNjbzEZMBcGA1UEChMQb3JnMS5leGFtcGxlLmNvbTEcMBoGA1UEAxMTY2Eu\nb3JnMS5leGFtcGxlLmNvbTAeFw0xOTA0MTYwNjExMDBaFw0yOTA0MTMwNjExMDBa\nMHMxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpDYWxpZm9ybmlhMRYwFAYDVQQHEw1T\nYW4gRnJhbmNpc2NvMRkwFwYDVQQKExBvcmcxLmV4YW1wbGUuY29tMRwwGgYDVQQD\nExNjYS5vcmcxLmV4YW1wbGUuY29tMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE\nL9sYtMP8suZ4Yu5rXn5mMPyoWtbGEZHLJZ5V0SxaSXY7pKfCSZHYRN0ItmQVjepc\nh8A6E9IVNPmN/lNq4Hq+c6NtMGswDgYDVR0PAQH/BAQDAgGmMB0GA1UdJQQWMBQG\nCCsGAQUFBwMCBggrBgEFBQcDATAPBgNVHRMBAf8EBTADAQH/MCkGA1UdDgQiBCCe\nNrgA+KqNCnTp9OGZEwHUWFNdgERzfZCfgl2y2hL1gDAKBggqhkjOPQQDAgNIADBF\nAiEAxeZDayv6lSf20PzDgfzP0naOvm2dpcDKkMg8eRrkFJMCIF2fI1XxSiZREfdR\nmJezwS1VfAagPjEivs2TTXvy7N6T\n-----END CERTIFICATE-----\n",
        //       tls_root_cert:
        //         "-----BEGIN CERTIFICATE-----\nMIICWDCCAf6gAwIBAgIRAKps39M67r2Lf/Eaob3Jon0wCgYIKoZIzj0EAwIwdjEL\nMAkGA1UEBhMCVVMxEzARBgNVBAgTCkNhbGlmb3JuaWExFjAUBgNVBAcTDVNhbiBG\ncmFuY2lzY28xGTAXBgNVBAoTEG9yZzEuZXhhbXBsZS5jb20xHzAdBgNVBAMTFnRs\nc2NhLm9yZzEuZXhhbXBsZS5jb20wHhcNMTkwNDE2MDYxMTAwWhcNMjkwNDEzMDYx\nMTAwWjB2MQswCQYDVQQGEwJVUzETMBEGA1UECBMKQ2FsaWZvcm5pYTEWMBQGA1UE\nBxMNU2FuIEZyYW5jaXNjbzEZMBcGA1UEChMQb3JnMS5leGFtcGxlLmNvbTEfMB0G\nA1UEAxMWdGxzY2Eub3JnMS5leGFtcGxlLmNvbTBZMBMGByqGSM49AgEGCCqGSM49\nAwEHA0IABA7b4/ZOsISIBP9JL5K4g0ItBk8NGS8LqF73aZ2QR943EtfmcfA4oA+H\nlwSdF6WIV4ZelH1NxoOOpV3KOw2ROrajbTBrMA4GA1UdDwEB/wQEAwIBpjAdBgNV\nHSUEFjAUBggrBgEFBQcDAgYIKwYBBQUHAwEwDwYDVR0TAQH/BAUwAwEB/zApBgNV\nHQ4EIgQgafK5Tg0xY5WFPkBFRk0K8hL5+7SKQNCfrUk4ZBsgatgwCgYIKoZIzj0E\nAwIDSAAwRQIhAKc9wwk2BWezACBW+cTm+Phl16OUs69nClMwNjuMutOpAiBmSUf3\nSpCPK5uzm0Ea5pJfWEARgiIUoIxnu73dVHCruw==\n-----END CERTIFICATE-----\n",
        //       admin:
        //         "-----BEGIN CERTIFICATE-----\nMIICKjCCAdGgAwIBAgIRAOy5I5ehAKNnpyfkonxDaHgwCgYIKoZIzj0EAwIwczEL\nMAkGA1UEBhMCVVMxEzARBgNVBAgTCkNhbGlmb3JuaWExFjAUBgNVBAcTDVNhbiBG\ncmFuY2lzY28xGTAXBgNVBAoTEG9yZzEuZXhhbXBsZS5jb20xHDAaBgNVBAMTE2Nh\nLm9yZzEuZXhhbXBsZS5jb20wHhcNMTkwNDE2MDYxMTAwWhcNMjkwNDEzMDYxMTAw\nWjBsMQswCQYDVQQGEwJVUzETMBEGA1UECBMKQ2FsaWZvcm5pYTEWMBQGA1UEBxMN\nU2FuIEZyYW5jaXNjbzEPMA0GA1UECxMGY2xpZW50MR8wHQYDVQQDDBZBZG1pbkBv\ncmcxLmV4YW1wbGUuY29tMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEgKR4o7+C\nKj6POyFWucZCDOwXUL/ggjL0nRLeMUmfE+6jidJlUv3QnDUSiENG7v3rjd6IT4QI\ngNLUAoqdW8fA3aNNMEswDgYDVR0PAQH/BAQDAgeAMAwGA1UdEwEB/wQCMAAwKwYD\nVR0jBCQwIoAgnja4APiqjQp06fThmRMB1FhTXYBEc32Qn4JdstoS9YAwCgYIKoZI\nzj0EAwIDRwAwRAIgSgjo9G1jhNee2LTtFP/xwSWwkTZ9IZ9WYCuPDty0TisCIBvs\nOzFtcyZ7acCQ+3NMwNEWvnoPclwgr8WOnj1Cxg42\n-----END CERTIFICATE-----\n"
        //     },
        //     {
        //       name: "Org2MSP",
        //       root_cert:
        //         "-----BEGIN CERTIFICATE-----\nMIICUTCCAfegAwIBAgIQMp4IBpoAWapZ39q29uz1DjAKBggqhkjOPQQDAjBzMQsw\nCQYDVQQGEwJVUzETMBEGA1UECBMKQ2FsaWZvcm5pYTEWMBQGA1UEBxMNU2FuIEZy\nYW5jaXNjbzEZMBcGA1UEChMQb3JnMi5leGFtcGxlLmNvbTEcMBoGA1UEAxMTY2Eu\nb3JnMi5leGFtcGxlLmNvbTAeFw0xOTA0MTYwNjExMDBaFw0yOTA0MTMwNjExMDBa\nMHMxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpDYWxpZm9ybmlhMRYwFAYDVQQHEw1T\nYW4gRnJhbmNpc2NvMRkwFwYDVQQKExBvcmcyLmV4YW1wbGUuY29tMRwwGgYDVQQD\nExNjYS5vcmcyLmV4YW1wbGUuY29tMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE\n/DshJB+zDPe19LA7MV7XXKooW/ZgY1oNuCfKUFW+I5PFR79y3yT46ZmH/aWD4q7B\nFcvh8SRIPdlmksF6vUYV+aNtMGswDgYDVR0PAQH/BAQDAgGmMB0GA1UdJQQWMBQG\nCCsGAQUFBwMCBggrBgEFBQcDATAPBgNVHRMBAf8EBTADAQH/MCkGA1UdDgQiBCA/\nOfMCj6ov/gW/sRXlOpVhQlQMZ7DGd4MTWcECYHdChzAKBggqhkjOPQQDAgNIADBF\nAiEA0/ATDSq8qX13DxpwMIfVpWv7kbP5kAIdzcKxVJwtaPECIG1hl5AesZrsUm4j\nnomzoy+9v7M4CW9wgUOgyGqNB1a9\n-----END CERTIFICATE-----\n",
        //       tls_root_cert:
        //         "-----BEGIN CERTIFICATE-----\nMIICVjCCAf2gAwIBAgIQQpLEihHohB8gpNtYO4qXhzAKBggqhkjOPQQDAjB2MQsw\nCQYDVQQGEwJVUzETMBEGA1UECBMKQ2FsaWZvcm5pYTEWMBQGA1UEBxMNU2FuIEZy\nYW5jaXNjbzEZMBcGA1UEChMQb3JnMi5leGFtcGxlLmNvbTEfMB0GA1UEAxMWdGxz\nY2Eub3JnMi5leGFtcGxlLmNvbTAeFw0xOTA0MTYwNjExMDBaFw0yOTA0MTMwNjEx\nMDBaMHYxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpDYWxpZm9ybmlhMRYwFAYDVQQH\nEw1TYW4gRnJhbmNpc2NvMRkwFwYDVQQKExBvcmcyLmV4YW1wbGUuY29tMR8wHQYD\nVQQDExZ0bHNjYS5vcmcyLmV4YW1wbGUuY29tMFkwEwYHKoZIzj0CAQYIKoZIzj0D\nAQcDQgAEQGDnPH+R7ocBNiKwcu1fFGoDSGj+j8x1qgKSUY9wZaFMk3LBO/Hf81Lj\no9Dm8JBPV0GKub9ayG5IHwpZZGJ9DKNtMGswDgYDVR0PAQH/BAQDAgGmMB0GA1Ud\nJQQWMBQGCCsGAQUFBwMCBggrBgEFBQcDATAPBgNVHRMBAf8EBTADAQH/MCkGA1Ud\nDgQiBCDZ8+z5+rFztvao6Yf5PELPg6JbJRU35kxPdyjVAgPawTAKBggqhkjOPQQD\nAgNHADBEAiBlxVdCIAO3yf73wB1yFYa3vAbvKZiv7bX+kGhLr+bLygIgWvCxnv7N\nHfGzlLaKmBwvAC6yzDIBofVvBE4XTa2mhCQ=\n-----END CERTIFICATE-----\n",
        //       admin:
        //         "-----BEGIN CERTIFICATE-----\nMIICKjCCAdGgAwIBAgIRAOZ4aVwEAISY9QP+RLLYukkwCgYIKoZIzj0EAwIwczEL\nMAkGA1UEBhMCVVMxEzARBgNVBAgTCkNhbGlmb3JuaWExFjAUBgNVBAcTDVNhbiBG\ncmFuY2lzY28xGTAXBgNVBAoTEG9yZzIuZXhhbXBsZS5jb20xHDAaBgNVBAMTE2Nh\nLm9yZzIuZXhhbXBsZS5jb20wHhcNMTkwNDE2MDYxMTAwWhcNMjkwNDEzMDYxMTAw\nWjBsMQswCQYDVQQGEwJVUzETMBEGA1UECBMKQ2FsaWZvcm5pYTEWMBQGA1UEBxMN\nU2FuIEZyYW5jaXNjbzEPMA0GA1UECxMGY2xpZW50MR8wHQYDVQQDDBZBZG1pbkBv\ncmcyLmV4YW1wbGUuY29tMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEIN5pAvKR\naq33jqbFzaHoDe7APsKGaNfUU2gbvVvYfRQIeI2hw/Ow+3RKvB1BfmqC6JrnsYMy\n0fr4PkmwE9eCj6NNMEswDgYDVR0PAQH/BAQDAgeAMAwGA1UdEwEB/wQCMAAwKwYD\nVR0jBCQwIoAgPznzAo+qL/4Fv7EV5TqVYUJUDGewxneDE1nBAmB3QocwCgYIKoZI\nzj0EAwIDRwAwRAIgR0y7shsNKob5Qn6/8iLeKN4CzCL15KrLAA75GKSVmpQCIElV\nV10pX+6IV9uZD7xrf2qqJUhzqwe+SOE/Q6YOi/M7\n-----END CERTIFICATE-----\n"
        //     }
        //   ],
        //   values: {
        //     consortium: "SampleConsortium",
        //     hashing_algorithm: "SHA256",
        //     orderer_addresses: [
        //       "orderer.example.com:7050",
        //       "orderer2.example.com:7050",
        //       "orderer3.example.com:7050",
        //       "orderer4.example.com:7050",
        //       "orderer5.example.com:7050"
        //     ]
        //   },
        //   consensus: {
        //     type: "etcdraft",
        //     raft_metadata: {
        //       consenters: [
        //         {
        //           host: "orderer.example.com",
        //           port: 7050,
        //           client_tls_cert:
        //             "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNXakNDQWdDZ0F3SUJBZ0lSQVBSVzAyYWZZNXpJRFQ4NG9IVVVGaWd3Q2dZSUtvWkl6ajBFQXdJd2JERUwKTUFrR0ExVUVCaE1DVlZNeEV6QVJCZ05WQkFnVENrTmhiR2xtYjNKdWFXRXhGakFVQmdOVkJBY1REVk5oYmlCRwpjbUZ1WTJselkyOHhGREFTQmdOVkJBb1RDMlY0WVcxd2JHVXVZMjl0TVJvd0dBWURWUVFERXhGMGJITmpZUzVsCmVHRnRjR3hsTG1OdmJUQWVGdzB4T1RBME1UWXdOakV4TURCYUZ3MHlPVEEwTVRNd05qRXhNREJhTUZneEN6QUoKQmdOVkJBWVRBbFZUTVJNd0VRWURWUVFJRXdwRFlXeHBabTl5Ym1saE1SWXdGQVlEVlFRSEV3MVRZVzRnUm5KaApibU5wYzJOdk1Sd3dHZ1lEVlFRREV4TnZjbVJsY21WeUxtVjRZVzF3YkdVdVkyOXRNRmt3RXdZSEtvWkl6ajBDCkFRWUlLb1pJemowREFRY0RRZ0FFZDJmMUVnbzlVZVFVSzZJdTEyN2Q5T1dWRmtxYmRHd2xFYWl4clkzNnd1YzMKS0dXRXZjUzhJVE53V252OC9QS3kvWU5JR0ZkMXZMUU1nbGxuMVhNcjBhT0JsakNCa3pBT0JnTlZIUThCQWY4RQpCQU1DQmFBd0hRWURWUjBsQkJZd0ZBWUlLd1lCQlFVSEF3RUdDQ3NHQVFVRkJ3TUNNQXdHQTFVZEV3RUIvd1FDCk1BQXdLd1lEVlIwakJDUXdJb0FneG1vMTd2Qzc1NUZWS3JJVnRFc0M2Q08rYzBodE5JelNHMjF1eEhwdUc0d3cKSndZRFZSMFJCQ0F3SG9JVGIzSmtaWEpsY2k1bGVHRnRjR3hsTG1OdmJZSUhiM0prWlhKbGNqQUtCZ2dxaGtqTwpQUVFEQWdOSUFEQkZBaUVBcHY2L0JMWG82Z3c1ZGgxTDFOUHBoakFtVnJaUm9ac0NCUjRnSGdlV21EZ0NJRDNTCjhPSlhtZjI0SVgrcUg4VEZjakhxT2MvOGNqYWpYTUZXUlNvL0E4N1cKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=",
        //           server_tls_cert:
        //             "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNXakNDQWdDZ0F3SUJBZ0lSQVBSVzAyYWZZNXpJRFQ4NG9IVVVGaWd3Q2dZSUtvWkl6ajBFQXdJd2JERUwKTUFrR0ExVUVCaE1DVlZNeEV6QVJCZ05WQkFnVENrTmhiR2xtYjNKdWFXRXhGakFVQmdOVkJBY1REVk5oYmlCRwpjbUZ1WTJselkyOHhGREFTQmdOVkJBb1RDMlY0WVcxd2JHVXVZMjl0TVJvd0dBWURWUVFERXhGMGJITmpZUzVsCmVHRnRjR3hsTG1OdmJUQWVGdzB4T1RBME1UWXdOakV4TURCYUZ3MHlPVEEwTVRNd05qRXhNREJhTUZneEN6QUoKQmdOVkJBWVRBbFZUTVJNd0VRWURWUVFJRXdwRFlXeHBabTl5Ym1saE1SWXdGQVlEVlFRSEV3MVRZVzRnUm5KaApibU5wYzJOdk1Sd3dHZ1lEVlFRREV4TnZjbVJsY21WeUxtVjRZVzF3YkdVdVkyOXRNRmt3RXdZSEtvWkl6ajBDCkFRWUlLb1pJemowREFRY0RRZ0FFZDJmMUVnbzlVZVFVSzZJdTEyN2Q5T1dWRmtxYmRHd2xFYWl4clkzNnd1YzMKS0dXRXZjUzhJVE53V252OC9QS3kvWU5JR0ZkMXZMUU1nbGxuMVhNcjBhT0JsakNCa3pBT0JnTlZIUThCQWY4RQpCQU1DQmFBd0hRWURWUjBsQkJZd0ZBWUlLd1lCQlFVSEF3RUdDQ3NHQVFVRkJ3TUNNQXdHQTFVZEV3RUIvd1FDCk1BQXdLd1lEVlIwakJDUXdJb0FneG1vMTd2Qzc1NUZWS3JJVnRFc0M2Q08rYzBodE5JelNHMjF1eEhwdUc0d3cKSndZRFZSMFJCQ0F3SG9JVGIzSmtaWEpsY2k1bGVHRnRjR3hsTG1OdmJZSUhiM0prWlhKbGNqQUtCZ2dxaGtqTwpQUVFEQWdOSUFEQkZBaUVBcHY2L0JMWG82Z3c1ZGgxTDFOUHBoakFtVnJaUm9ac0NCUjRnSGdlV21EZ0NJRDNTCjhPSlhtZjI0SVgrcUg4VEZjakhxT2MvOGNqYWpYTUZXUlNvL0E4N1cKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
        //         },
        //         {
        //           host: "orderer2.example.com",
        //           port: 7050,
        //           client_tls_cert:
        //             "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNYRENDQWdLZ0F3SUJBZ0lRWkVEQklSYUNiMmJMVXA3ZXMrZEJVVEFLQmdncWhrak9QUVFEQWpCc01Rc3cKQ1FZRFZRUUdFd0pWVXpFVE1CRUdBMVVFQ0JNS1EyRnNhV1p2Y201cFlURVdNQlFHQTFVRUJ4TU5VMkZ1SUVaeQpZVzVqYVhOamJ6RVVNQklHQTFVRUNoTUxaWGhoYlhCc1pTNWpiMjB4R2pBWUJnTlZCQU1URVhSc2MyTmhMbVY0CllXMXdiR1V1WTI5dE1CNFhEVEU1TURReE5qQTJNVEV3TUZvWERUSTVNRFF4TXpBMk1URXdNRm93V1RFTE1Ba0cKQTFVRUJoTUNWVk14RXpBUkJnTlZCQWdUQ2tOaGJHbG1iM0p1YVdFeEZqQVVCZ05WQkFjVERWTmhiaUJHY21GdQpZMmx6WTI4eEhUQWJCZ05WQkFNVEZHOXlaR1Z5WlhJeUxtVjRZVzF3YkdVdVkyOXRNRmt3RXdZSEtvWkl6ajBDCkFRWUlLb1pJemowREFRY0RRZ0FFZnpxUk9HMEg1UEtzckJUbzFJdFNmanN3bVp3emplZnZpa1BwUDRaTm9XNVYKYTRKUFNpaTJ6TEgwMUNoUXY0b3Q0V2hQNHpSV251NHIzOG02RWwrNSthT0JtRENCbFRBT0JnTlZIUThCQWY4RQpCQU1DQmFBd0hRWURWUjBsQkJZd0ZBWUlLd1lCQlFVSEF3RUdDQ3NHQVFVRkJ3TUNNQXdHQTFVZEV3RUIvd1FDCk1BQXdLd1lEVlIwakJDUXdJb0FneG1vMTd2Qzc1NUZWS3JJVnRFc0M2Q08rYzBodE5JelNHMjF1eEhwdUc0d3cKS1FZRFZSMFJCQ0l3SUlJVWIzSmtaWEpsY2pJdVpYaGhiWEJzWlM1amIyMkNDRzl5WkdWeVpYSXlNQW9HQ0NxRwpTTTQ5QkFNQ0EwZ0FNRVVDSVFEaHN5S1VCL3J4N1I2b2pNSTRNbVVTUGdQbXBRVTVicWRZRUQzbFdUc2F5UUlnCkFuVE9yQTBZTHFPWjF2eklwSnVDaStYeEhEN3FtT1A3SFBDdEs0UG00RGM9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K",
        //           server_tls_cert:
        //             "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNYRENDQWdLZ0F3SUJBZ0lRWkVEQklSYUNiMmJMVXA3ZXMrZEJVVEFLQmdncWhrak9QUVFEQWpCc01Rc3cKQ1FZRFZRUUdFd0pWVXpFVE1CRUdBMVVFQ0JNS1EyRnNhV1p2Y201cFlURVdNQlFHQTFVRUJ4TU5VMkZ1SUVaeQpZVzVqYVhOamJ6RVVNQklHQTFVRUNoTUxaWGhoYlhCc1pTNWpiMjB4R2pBWUJnTlZCQU1URVhSc2MyTmhMbVY0CllXMXdiR1V1WTI5dE1CNFhEVEU1TURReE5qQTJNVEV3TUZvWERUSTVNRFF4TXpBMk1URXdNRm93V1RFTE1Ba0cKQTFVRUJoTUNWVk14RXpBUkJnTlZCQWdUQ2tOaGJHbG1iM0p1YVdFeEZqQVVCZ05WQkFjVERWTmhiaUJHY21GdQpZMmx6WTI4eEhUQWJCZ05WQkFNVEZHOXlaR1Z5WlhJeUxtVjRZVzF3YkdVdVkyOXRNRmt3RXdZSEtvWkl6ajBDCkFRWUlLb1pJemowREFRY0RRZ0FFZnpxUk9HMEg1UEtzckJUbzFJdFNmanN3bVp3emplZnZpa1BwUDRaTm9XNVYKYTRKUFNpaTJ6TEgwMUNoUXY0b3Q0V2hQNHpSV251NHIzOG02RWwrNSthT0JtRENCbFRBT0JnTlZIUThCQWY4RQpCQU1DQmFBd0hRWURWUjBsQkJZd0ZBWUlLd1lCQlFVSEF3RUdDQ3NHQVFVRkJ3TUNNQXdHQTFVZEV3RUIvd1FDCk1BQXdLd1lEVlIwakJDUXdJb0FneG1vMTd2Qzc1NUZWS3JJVnRFc0M2Q08rYzBodE5JelNHMjF1eEhwdUc0d3cKS1FZRFZSMFJCQ0l3SUlJVWIzSmtaWEpsY2pJdVpYaGhiWEJzWlM1amIyMkNDRzl5WkdWeVpYSXlNQW9HQ0NxRwpTTTQ5QkFNQ0EwZ0FNRVVDSVFEaHN5S1VCL3J4N1I2b2pNSTRNbVVTUGdQbXBRVTVicWRZRUQzbFdUc2F5UUlnCkFuVE9yQTBZTHFPWjF2eklwSnVDaStYeEhEN3FtT1A3SFBDdEs0UG00RGM9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"
        //         },
        //         {
        //           host: "orderer3.example.com",
        //           port: 7050,
        //           client_tls_cert:
        //             "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNXekNDQWdLZ0F3SUJBZ0lRSnZ5dHlyWVUwMVpTbmdaTGxaNFRtekFLQmdncWhrak9QUVFEQWpCc01Rc3cKQ1FZRFZRUUdFd0pWVXpFVE1CRUdBMVVFQ0JNS1EyRnNhV1p2Y201cFlURVdNQlFHQTFVRUJ4TU5VMkZ1SUVaeQpZVzVqYVhOamJ6RVVNQklHQTFVRUNoTUxaWGhoYlhCc1pTNWpiMjB4R2pBWUJnTlZCQU1URVhSc2MyTmhMbVY0CllXMXdiR1V1WTI5dE1CNFhEVEU1TURReE5qQTJNVEV3TUZvWERUSTVNRFF4TXpBMk1URXdNRm93V1RFTE1Ba0cKQTFVRUJoTUNWVk14RXpBUkJnTlZCQWdUQ2tOaGJHbG1iM0p1YVdFeEZqQVVCZ05WQkFjVERWTmhiaUJHY21GdQpZMmx6WTI4eEhUQWJCZ05WQkFNVEZHOXlaR1Z5WlhJekxtVjRZVzF3YkdVdVkyOXRNRmt3RXdZSEtvWkl6ajBDCkFRWUlLb1pJemowREFRY0RRZ0FFYmUxaUhRT0NxZlYyWjRMZ2ZxcnJSWFdXL0labzZrSXJDbWdSY0tnclNVbmwKaGlra2lsNEpvT2NWNnc3eXlVcXhnZDNtamlTcXV6RlYvU0RwTWpwR1dxT0JtRENCbFRBT0JnTlZIUThCQWY4RQpCQU1DQmFBd0hRWURWUjBsQkJZd0ZBWUlLd1lCQlFVSEF3RUdDQ3NHQVFVRkJ3TUNNQXdHQTFVZEV3RUIvd1FDCk1BQXdLd1lEVlIwakJDUXdJb0FneG1vMTd2Qzc1NUZWS3JJVnRFc0M2Q08rYzBodE5JelNHMjF1eEhwdUc0d3cKS1FZRFZSMFJCQ0l3SUlJVWIzSmtaWEpsY2pNdVpYaGhiWEJzWlM1amIyMkNDRzl5WkdWeVpYSXpNQW9HQ0NxRwpTTTQ5QkFNQ0EwY0FNRVFDSUVXTWNXcTRqZkkzWXp3d0VSWkpjUzJnOVVPSHBneEpOS2FTQnNseGh1ZTRBaUJJCmlPQnZFWmxnKzJ3anU1NUxqNUJiT1ZpUU8vUkZZNXlGajl0Sm1vRVZ0QT09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K",
        //           server_tls_cert:
        //             "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNXekNDQWdLZ0F3SUJBZ0lRSnZ5dHlyWVUwMVpTbmdaTGxaNFRtekFLQmdncWhrak9QUVFEQWpCc01Rc3cKQ1FZRFZRUUdFd0pWVXpFVE1CRUdBMVVFQ0JNS1EyRnNhV1p2Y201cFlURVdNQlFHQTFVRUJ4TU5VMkZ1SUVaeQpZVzVqYVhOamJ6RVVNQklHQTFVRUNoTUxaWGhoYlhCc1pTNWpiMjB4R2pBWUJnTlZCQU1URVhSc2MyTmhMbVY0CllXMXdiR1V1WTI5dE1CNFhEVEU1TURReE5qQTJNVEV3TUZvWERUSTVNRFF4TXpBMk1URXdNRm93V1RFTE1Ba0cKQTFVRUJoTUNWVk14RXpBUkJnTlZCQWdUQ2tOaGJHbG1iM0p1YVdFeEZqQVVCZ05WQkFjVERWTmhiaUJHY21GdQpZMmx6WTI4eEhUQWJCZ05WQkFNVEZHOXlaR1Z5WlhJekxtVjRZVzF3YkdVdVkyOXRNRmt3RXdZSEtvWkl6ajBDCkFRWUlLb1pJemowREFRY0RRZ0FFYmUxaUhRT0NxZlYyWjRMZ2ZxcnJSWFdXL0labzZrSXJDbWdSY0tnclNVbmwKaGlra2lsNEpvT2NWNnc3eXlVcXhnZDNtamlTcXV6RlYvU0RwTWpwR1dxT0JtRENCbFRBT0JnTlZIUThCQWY4RQpCQU1DQmFBd0hRWURWUjBsQkJZd0ZBWUlLd1lCQlFVSEF3RUdDQ3NHQVFVRkJ3TUNNQXdHQTFVZEV3RUIvd1FDCk1BQXdLd1lEVlIwakJDUXdJb0FneG1vMTd2Qzc1NUZWS3JJVnRFc0M2Q08rYzBodE5JelNHMjF1eEhwdUc0d3cKS1FZRFZSMFJCQ0l3SUlJVWIzSmtaWEpsY2pNdVpYaGhiWEJzWlM1amIyMkNDRzl5WkdWeVpYSXpNQW9HQ0NxRwpTTTQ5QkFNQ0EwY0FNRVFDSUVXTWNXcTRqZkkzWXp3d0VSWkpjUzJnOVVPSHBneEpOS2FTQnNseGh1ZTRBaUJJCmlPQnZFWmxnKzJ3anU1NUxqNUJiT1ZpUU8vUkZZNXlGajl0Sm1vRVZ0QT09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"
        //         },
        //         {
        //           host: "orderer4.example.com",
        //           port: 7050,
        //           client_tls_cert:
        //             "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNYVENDQWdPZ0F3SUJBZ0lSQU5lc2JTR1Y2RWo5QzJ0RGVMN0gxaGN3Q2dZSUtvWkl6ajBFQXdJd2JERUwKTUFrR0ExVUVCaE1DVlZNeEV6QVJCZ05WQkFnVENrTmhiR2xtYjNKdWFXRXhGakFVQmdOVkJBY1REVk5oYmlCRwpjbUZ1WTJselkyOHhGREFTQmdOVkJBb1RDMlY0WVcxd2JHVXVZMjl0TVJvd0dBWURWUVFERXhGMGJITmpZUzVsCmVHRnRjR3hsTG1OdmJUQWVGdzB4T1RBME1UWXdOakV4TURCYUZ3MHlPVEEwTVRNd05qRXhNREJhTUZreEN6QUoKQmdOVkJBWVRBbFZUTVJNd0VRWURWUVFJRXdwRFlXeHBabTl5Ym1saE1SWXdGQVlEVlFRSEV3MVRZVzRnUm5KaApibU5wYzJOdk1SMHdHd1lEVlFRREV4UnZjbVJsY21WeU5DNWxlR0Z0Y0d4bExtTnZiVEJaTUJNR0J5cUdTTTQ5CkFnRUdDQ3FHU000OUF3RUhBMElBQkZhQks5cTk3Yzd1ZDR1TG41bHY5ZEFxcWFYZkZIeUl1M2I3aE95b0FuY2kKdVowUHIzb0ZwMVFjaDRiYlpQWlVvbWJnSWNxYVlDQzI4QzRzVjhhRk41S2pnWmd3Z1pVd0RnWURWUjBQQVFILwpCQVFEQWdXZ01CMEdBMVVkSlFRV01CUUdDQ3NHQVFVRkJ3TUJCZ2dyQmdFRkJRY0RBakFNQmdOVkhSTUJBZjhFCkFqQUFNQ3NHQTFVZEl3UWtNQ0tBSU1acU5lN3d1K2VSVlNxeUZiUkxBdWdqdm5OSWJUU00waHR0YnNSNmJodU0KTUNrR0ExVWRFUVFpTUNDQ0ZHOXlaR1Z5WlhJMExtVjRZVzF3YkdVdVkyOXRnZ2h2Y21SbGNtVnlOREFLQmdncQpoa2pPUFFRREFnTklBREJGQWlFQXNTYzJKbVBnTTZva1hSNVE5aG10Z3lZbTZUK2xiWGpGL1g2NlNrMGZVR29DCklGT2ZMbG9BY2FaYWVkSVY1V3J3dHFjandSUU00QXVETndjRzZlSGJsdnBvCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K",
        //           server_tls_cert:
        //             "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNYVENDQWdPZ0F3SUJBZ0lSQU5lc2JTR1Y2RWo5QzJ0RGVMN0gxaGN3Q2dZSUtvWkl6ajBFQXdJd2JERUwKTUFrR0ExVUVCaE1DVlZNeEV6QVJCZ05WQkFnVENrTmhiR2xtYjNKdWFXRXhGakFVQmdOVkJBY1REVk5oYmlCRwpjbUZ1WTJselkyOHhGREFTQmdOVkJBb1RDMlY0WVcxd2JHVXVZMjl0TVJvd0dBWURWUVFERXhGMGJITmpZUzVsCmVHRnRjR3hsTG1OdmJUQWVGdzB4T1RBME1UWXdOakV4TURCYUZ3MHlPVEEwTVRNd05qRXhNREJhTUZreEN6QUoKQmdOVkJBWVRBbFZUTVJNd0VRWURWUVFJRXdwRFlXeHBabTl5Ym1saE1SWXdGQVlEVlFRSEV3MVRZVzRnUm5KaApibU5wYzJOdk1SMHdHd1lEVlFRREV4UnZjbVJsY21WeU5DNWxlR0Z0Y0d4bExtTnZiVEJaTUJNR0J5cUdTTTQ5CkFnRUdDQ3FHU000OUF3RUhBMElBQkZhQks5cTk3Yzd1ZDR1TG41bHY5ZEFxcWFYZkZIeUl1M2I3aE95b0FuY2kKdVowUHIzb0ZwMVFjaDRiYlpQWlVvbWJnSWNxYVlDQzI4QzRzVjhhRk41S2pnWmd3Z1pVd0RnWURWUjBQQVFILwpCQVFEQWdXZ01CMEdBMVVkSlFRV01CUUdDQ3NHQVFVRkJ3TUJCZ2dyQmdFRkJRY0RBakFNQmdOVkhSTUJBZjhFCkFqQUFNQ3NHQTFVZEl3UWtNQ0tBSU1acU5lN3d1K2VSVlNxeUZiUkxBdWdqdm5OSWJUU00waHR0YnNSNmJodU0KTUNrR0ExVWRFUVFpTUNDQ0ZHOXlaR1Z5WlhJMExtVjRZVzF3YkdVdVkyOXRnZ2h2Y21SbGNtVnlOREFLQmdncQpoa2pPUFFRREFnTklBREJGQWlFQXNTYzJKbVBnTTZva1hSNVE5aG10Z3lZbTZUK2xiWGpGL1g2NlNrMGZVR29DCklGT2ZMbG9BY2FaYWVkSVY1V3J3dHFjandSUU00QXVETndjRzZlSGJsdnBvCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"
        //         },
        //         {
        //           host: "orderer5.example.com",
        //           port: 7050,
        //           client_tls_cert:
        //             "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNYRENDQWdPZ0F3SUJBZ0lSQUtzQkd0Y2VIMlFocjdlNjAyaGpncWt3Q2dZSUtvWkl6ajBFQXdJd2JERUwKTUFrR0ExVUVCaE1DVlZNeEV6QVJCZ05WQkFnVENrTmhiR2xtYjNKdWFXRXhGakFVQmdOVkJBY1REVk5oYmlCRwpjbUZ1WTJselkyOHhGREFTQmdOVkJBb1RDMlY0WVcxd2JHVXVZMjl0TVJvd0dBWURWUVFERXhGMGJITmpZUzVsCmVHRnRjR3hsTG1OdmJUQWVGdzB4T1RBME1UWXdOakV4TURCYUZ3MHlPVEEwTVRNd05qRXhNREJhTUZreEN6QUoKQmdOVkJBWVRBbFZUTVJNd0VRWURWUVFJRXdwRFlXeHBabTl5Ym1saE1SWXdGQVlEVlFRSEV3MVRZVzRnUm5KaApibU5wYzJOdk1SMHdHd1lEVlFRREV4UnZjbVJsY21WeU5TNWxlR0Z0Y0d4bExtTnZiVEJaTUJNR0J5cUdTTTQ5CkFnRUdDQ3FHU000OUF3RUhBMElBQklWVVBjeXBIdlE1M05sOTE2NytSdDZkeERrSmhXNGVHdGZwcWhmSjZ5eWcKOXZKczgzVmRqdm55WHBPcE15dENBblJFNGk3YmtvbkpJMXEyc2pkaUt3MmpnWmd3Z1pVd0RnWURWUjBQQVFILwpCQVFEQWdXZ01CMEdBMVVkSlFRV01CUUdDQ3NHQVFVRkJ3TUJCZ2dyQmdFRkJRY0RBakFNQmdOVkhSTUJBZjhFCkFqQUFNQ3NHQTFVZEl3UWtNQ0tBSU1acU5lN3d1K2VSVlNxeUZiUkxBdWdqdm5OSWJUU00waHR0YnNSNmJodU0KTUNrR0ExVWRFUVFpTUNDQ0ZHOXlaR1Z5WlhJMUxtVjRZVzF3YkdVdVkyOXRnZ2h2Y21SbGNtVnlOVEFLQmdncQpoa2pPUFFRREFnTkhBREJFQWlCY3Jpay9GdjVRaVlPTmtHNE5DZE43a0VsT3dZZHIwUXR5MXFqSlQ2T3ZmUUlnCkhNdUpINHBacjJNVk1EcWdzWmpFaGtQZU1HZUovSjB5SVU2M0Jib2lXZmc9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K",
        //           server_tls_cert:
        //             "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNYRENDQWdPZ0F3SUJBZ0lSQUtzQkd0Y2VIMlFocjdlNjAyaGpncWt3Q2dZSUtvWkl6ajBFQXdJd2JERUwKTUFrR0ExVUVCaE1DVlZNeEV6QVJCZ05WQkFnVENrTmhiR2xtYjNKdWFXRXhGakFVQmdOVkJBY1REVk5oYmlCRwpjbUZ1WTJselkyOHhGREFTQmdOVkJBb1RDMlY0WVcxd2JHVXVZMjl0TVJvd0dBWURWUVFERXhGMGJITmpZUzVsCmVHRnRjR3hsTG1OdmJUQWVGdzB4T1RBME1UWXdOakV4TURCYUZ3MHlPVEEwTVRNd05qRXhNREJhTUZreEN6QUoKQmdOVkJBWVRBbFZUTVJNd0VRWURWUVFJRXdwRFlXeHBabTl5Ym1saE1SWXdGQVlEVlFRSEV3MVRZVzRnUm5KaApibU5wYzJOdk1SMHdHd1lEVlFRREV4UnZjbVJsY21WeU5TNWxlR0Z0Y0d4bExtTnZiVEJaTUJNR0J5cUdTTTQ5CkFnRUdDQ3FHU000OUF3RUhBMElBQklWVVBjeXBIdlE1M05sOTE2NytSdDZkeERrSmhXNGVHdGZwcWhmSjZ5eWcKOXZKczgzVmRqdm55WHBPcE15dENBblJFNGk3YmtvbkpJMXEyc2pkaUt3MmpnWmd3Z1pVd0RnWURWUjBQQVFILwpCQVFEQWdXZ01CMEdBMVVkSlFRV01CUUdDQ3NHQVFVRkJ3TUJCZ2dyQmdFRkJRY0RBakFNQmdOVkhSTUJBZjhFCkFqQUFNQ3NHQTFVZEl3UWtNQ0tBSU1acU5lN3d1K2VSVlNxeUZiUkxBdWdqdm5OSWJUU00waHR0YnNSNmJodU0KTUNrR0ExVWRFUVFpTUNDQ0ZHOXlaR1Z5WlhJMUxtVjRZVzF3YkdVdVkyOXRnZ2h2Y21SbGNtVnlOVEFLQmdncQpoa2pPUFFRREFnTkhBREJFQWlCY3Jpay9GdjVRaVlPTmtHNE5DZE43a0VsT3dZZHIwUXR5MXFqSlQ2T3ZmUUlnCkhNdUpINHBacjJNVk1EcWdzWmpFaGtQZU1HZUovSjB5SVU2M0Jib2lXZmc9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"
        //         }
        //       ],
        //       options: {
        //         tick_interval: "500ms",
        //         election_tick: 10,
        //         heartbeat_tick: 1,
        //         max_inflight_blocks: 5,
        //         snapshot_interval_size: 20971520
        //       }
        //     },
        //     max_message_count: 10,
        //     absolute_max_bytes: 103809024,
        //     preferred_max_bytes: 524288,
        //     batch_time_out: "2s"
        //   }
        // }
      },
      consortium: false
    };
  },
  methods: {
    init() {
      if (localStorage.block) {
        this.block = JSON.parse(localStorage.block);
      } else {
        this.$message({
          message: "没有发现区块数据，请重新上传",
          type: "warnning"
        });
      }
      // console.log(this.block);
      // console.log(this.block.config.consortium_orgs);
    }
  },
  mounted() {
    this.init();
  }
};

function show(e) {
  console.log("show");
}
</script>

<style lang="less">
.box-card {
  margin-top: 5px;
}

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

.org_list {
  white-space: nowrap;
  list-style: none;
  overflow-x: scroll;
  text-align: left;
}

.org {
  display: inline-block;
  margin-left: 20px;
}
</style>
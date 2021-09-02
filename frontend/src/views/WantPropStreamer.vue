<template>
  <v-container
    fluid
    style="
      width: 50%;
      display: flex;
      align-items: center;
      justify-content: flex-start;
      flex-direction: column;
      margin: 0 auto;
    "
  >
    <v-row>
      <v-col cols="6"><!-- item-text will change-->
        <v-combobox
          style="width: 100%"
          v-model="selectedSeller"
          @change="takeSellerUsername()"
          :items="wantedUsers"
          label="Buyer Username"
          item-text="buyerUsername" 
          item-value="buyerUsername"
          clearable small-chips outlined dense
        />
      </v-col>
      <v-col cols="6"><!-- item-text will change-->
        <v-combobox
          style="width: 100%"
          v-model="selectedUnique"
          :items="sellerusername"
          label="Buyed Package"
          item-text="unique" 
          item-value="unique"
          clearable small-chips outlined dense
        />
      </v-col>
       <v-col cols="6">
         Completed
        <v-card
          class="mx-auto"
          max-width="800"
          tile
          v-for="item in storedPlan" :key="item.package.unique"
          >
          <div v-if="selectedUnique.unique==item.package.unique">
            <div v-for="item2 in item.package.items" :key="item2.prop">
              <v-list-item three-line v-if="item2.status=='completed'">
                  <v-list-item-content>
                      <v-list-item-title>{{item2.prop}}</v-list-item-title>
                      <v-list-item-action> 
                        <v-btn @click="wantProp(item.package.unique,item2.prop,'')" >Get Back</v-btn>
                      </v-list-item-action>
                  </v-list-item-content>
              </v-list-item>
            </div>
          </div>
        </v-card>
      </v-col>
       <v-col cols="6" >
         Requested
        <v-card
          class="mx-auto"
          max-width="800"
          tile
          v-for="item in storedPlan" :key="item.package.unique"
          >
          <div v-if="selectedUnique.unique==item.package.unique">
            <div v-for="item2 in item.package.items" :key="item2.prop">
              <v-list-item three-line v-if="item2.status=='requested'">
                  <v-list-item-content>
                      <v-list-item-title>{{item2.prop}}</v-list-item-title>
                      <v-list-item-action> 
                        <v-btn @click="wantProp(item.package.unique,item2.prop,'complete')" >Complete</v-btn>
                      </v-list-item-action>
                  </v-list-item-content>
              </v-list-item>
            </div>
          </div>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import axios from "axios";
import { mapState } from "vuex";
export default {
  name: "WantPropStreamer",
  computed: {
    ...mapState(["username", "host","type"]),
  },
  data() {
    return {
      success:"",
      sellerusername:[],
      storedPlan:[],
      wantedUsers:[],
      selectedSeller: {buyerUsername:"Choose"},
      selectedUnique: {unique:"Choose"}
    };
  },
  methods: {
    wantProp(uniqueCame,propCame,statusCame){
      axios
      .post(this.host+"/create/wantStreamer",{
        unique:uniqueCame,
        prop:propCame,
        buyerUsername:this.selectedSeller,
        sellerUsername:this.username,
        status:statusCame,
        output:"json"
      })
      .then(()=>{
        if (statusCame=="complete"){
          statusCame="completed"
        }
        for(let data in this.storedPlan){
          if(this.selectedUnique.unique==this.storedPlan[data].package.unique){
            for(let data2 in this.storedPlan[data].package.items){
              if(propCame==this.storedPlan[data].package.items[data2].prop){
                this.storedPlan[data].package.items[data2].status=statusCame
              }
            }
          }
        }
      })
      .catch((error) => {
          window.alert(`The API returned an error: ${error.response.data}`);
      });
    },
    takeSellerUsername() {
      this.sellerusername=[]//for delete datas in combobox
      for(let prop in this.storedPlan){
        for(let prop2 in this.storedPlan[prop].package.items){
          if (this.selectedSeller == this.storedPlan[prop].package.items[prop2].buyerUsername){
              this.sellerusername=[...this.sellerusername,this.storedPlan[prop].package]
          }
        }
      }
    },
    takeWantedUsers(){
      for(let prop in this.storedPlan){
        for(let prop2 in this.storedPlan[prop].package.items){
          this.wantedUsers=[...this.wantedUsers,this.storedPlan[prop].package.items[prop2].buyerUsername]
        }
      }
    },
    getStocks() {
      axios
        .put(this.host + "/list/userPlan", {
          username: this.username,
        }
      )
        .then((response) => {
          this.storedPlan=response.data[0].plan;
          this.takeWantedUsers();
        }
      );
    },
  },
  created() {
      this.getStocks();
  },
};
</script>

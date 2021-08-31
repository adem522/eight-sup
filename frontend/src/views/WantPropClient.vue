<template>
  <v-container
    fluid
    style="
      width: 65%;
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
          :items="storedPlan"
          label="Seller Username"
          item-text="sellerusername" 
          item-value="sellerusername"
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
      <v-col cols="4" >
         Completed
        <v-card
          class="mx-auto"
          max-width="800"
          tile
          v-for="item in storedPlan" :key="item.package.unique"
          >
          <div v-if="selectedUnique.unique==item.package.unique">
            <div v-for="item2 in item.package.items" :key="item2.prop">
              <v-list-item v-if="item2.status=='completed'" >
                  <v-list-item-content >
                      <v-list-item-title>{{item2.prop}}</v-list-item-title>
                      <v-list-item-action> 
                        <v-btn @click="wantProp(item.package.unique,item2.prop,'delete')">Delete</v-btn>
                      </v-list-item-action>
                  </v-list-item-content>
              </v-list-item>
            </div>
          </div>
        </v-card>
      </v-col>
      <v-col cols="4" >
         Requested
        <v-card
          class="mx-auto"
          max-width="800"
          tile
          v-for="item in storedPlan" :key="item.package.unique"
          >
          <div v-if="selectedUnique.unique==item.package.unique">
            <div v-for="item2 in item.package.items" :key="item2.prop">
              <v-list-item v-if="item2.status=='requested'" >
                  <v-list-item-content >
                      <v-list-item-title>{{item2.prop}}</v-list-item-title>
                      <v-list-item-action> 
                        <v-btn @click="wantProp(item.package.unique,item2.prop,'getBack')" >Get Back</v-btn>
                      </v-list-item-action>
                  </v-list-item-content>
              </v-list-item>
            </div>
          </div>
        </v-card>
      </v-col>
       <v-col cols="4" >
         Available
        <v-card
          class="mx-auto"
          max-width="800"
          tile
          v-for="item in storedPlan" :key="item.package.unique"
          >
          <div v-if="selectedUnique.unique==item.package.unique && selectedSeller.sellerusername==item.sellerusername">
            <div v-for="item2 in item.package.items" :key="item2.prop">
              <v-list-item v-if="item2.status=='available'" >
                  <v-list-item-content >
                      <v-list-item-title>{{item2.prop}}</v-list-item-title>
                      <v-list-item-title>{{item.package.unique}}</v-list-item-title>
                      <v-list-item-action> 
                        <v-btn @click="wantProp(item.package.unique,item2.prop,'want')" >Want</v-btn>
                      </v-list-item-action>
                  </v-list-item-content>
              </v-list-item>
            </div>
          </div>
        </v-card>
      </v-col>
      <v-col v-if="!!success" cols="12">
        <v-alert type="success"> {{success}} </v-alert>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import axios from "axios";
import { mapState } from "vuex";
export default {
  name: "WantPropClient",
  computed: {
    ...mapState(["username", "host","type","storedInfo"]),
  },
  data() {
    return {
      success:"",
      sellerusername:[],
      storedPlan:[],
      unique:[],
      selectedSeller: {sellerusername:"Choose"},
      selectedUnique: {unique:"Choose"}
    };
  },
  methods: {
    wantProp(uniqueCame,propCame,statusCame){
      axios
      .post(this.host+"/create/wantClient",{
        unique:uniqueCame,
        prop:propCame,
        status:statusCame,
        buyerUsername:this.username,
        sellerUsername:this.selectedSeller.sellerusername,
        output:"json"
      })
      .then((res)=>{
        this.success=res.data[0]
        this.hide_alert()
        if (statusCame=="want"){
          statusCame="requested"
        }
        else if (statusCame=="getBack"){
          statusCame="available"
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
          window.alert(`The API returned an error: ${error.response}`);
      });
    },
    takeSellerUsername() {
      this.sellerusername=[]//for delete datas in combobox
      this.unique=[]        //for delete datas in combobox
      for(let prop in this.storedPlan){
          if (this.selectedSeller.sellerusername == this.storedPlan[prop].sellerusername){
            this.sellerusername=[...this.sellerusername,this.storedPlan[prop].package]
        }
      }
    },
    getPlans() {
      axios
        .put(this.host + "/list/userPlan", {
          username: this.username,
        }
      )
        .then((response) => {
          this.storedPlan=response.data[0].plan;
          console.log(this.storedPlan)
        }
      );
    },
    getInfo() {
      if(this.storedInfo.length==0){
        axios
          .get(this.host + "/list/planInfo")
          .then((response) => {
            this.$store.commit("setStoredInfo", response.data.data);
        });
      }
    },
    getWants() {
      axios
        .put(this.host + "/list/userStock", {
          username: this.username,
        }
      )
        .then((response) => {
          this.storedPlan=response.data[0].plan;
        })
        .catch((error) => {
          window.alert(`The API returned an error: ${error.response.data}`);
        });
    },
    hide_alert() {
      window.setInterval(() => {
        this.success="";
      }, 1500)    
    }
  },
  created() {
      this.getPlans();
      this.getInfo();
  },
};
</script>

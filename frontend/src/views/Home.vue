<template>
  <div>
    <v-card
        class="mx-auto"
    >
    <div class="d-flex flex-no-wrap justify-space-between">
      <h2>Package Names</h2>
    </div>

      <v-container>
        <v-row dense>
          <v-col
            v-for="(item, i) in storedInfo"
            :key="i"
            cols="2"
            >
            <v-card
                :color="item.color"
                dark
            >
              <div class="d-flex flex-no-wrap justify-space-between">
                <div>
                  <v-card-title
                    class="text-h5"
                    v-text="item.name"/>
                  <v-card-subtitle 
                    v-text="'Cost '+item.cost+'$'"/>
                  <v-card-subtitle 
                    v-for="item2 in item.items" :key="item2" 
                    v-text="item2"/>
                </div>
              </div>
              <v-avatar>
                <v-img :src="item.img"/>
              </v-avatar>
            </v-card>
          </v-col>
        </v-row>
      </v-container>
    </v-card>
  </div>
</template>

<script>
import axios from "axios";
import { mapState } from "vuex";
  export default {
    name:"Home",
    computed:{
      ...mapState(["host","storedInfo"])
    },
    data: () => ({
      stored:[],
      items: [
        {
          //bronze
          img:'https://www.dupageseniorcouncil.org/wp-content/uploads/2018/02/Bronze-Icon-1.jpg',
        },
        {
          //silver
          img:'http://www.tpsnewengland.com/wp-content/uploads/2017/05/Silver-Icon.png',
        },
        {
          //gold
          img:'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQrhTfvz6h6Ok75qHkPdTQIwmFZLmiFvu_-MA&usqp=CAU',
        },
        {
          //emerald
          img:'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR6zeMLiTquLZe3z0YzqoH1rHhrv9m0t5mbDw&usqp=CAU'
        },
        {
          //vibranium
          img:'https://m.media-amazon.com/images/I/71ZfjP-Ui-L._AC_UL1500_.jpg',
        },
        {
          //diamond
          img:'https://thumbs.dreamstime.com/b/geometric-diamond-icon-low-poly-55394702.jpg',
        },
      ],
    }),
    methods:{
      getInfo() {
        axios
        .get(this.host + "/list/planInfo")
        .then((response) => {
          this.stored=response.data.data
            localStorage.setItem("storedInfo", response.data.data);
            this.$store.commit("setStoredInfo", response.data.data);
          for(let prop in this.storedInfo){
            this.storedInfo[prop].img=this.items[prop].img
          }
        });
      },
    },
    created() {
      this.getInfo();
    },
  }
</script>
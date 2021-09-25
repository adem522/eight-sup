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
          v-model="selectedBuyer"
          @change="takeBuyerUsername()"
          :items="storedEvent"
          label="Buyer Username"
          item-text="buyerUsername" 
          item-value="buyerUsername"
          clearable small-chips outlined dense
        />
      </v-col>
       <v-col cols="6">
        <v-card
          class="mx-auto"
          max-width="800"
          tile
          v-for="item in buyerusername" :key="item.unique"
          >
          <v-list-item three-line>
              <v-list-item-content>
                  <v-list-item-title >{{item.unique}}</v-list-item-title>
                  <v-list-item-subtitle>{{item.date}}</v-list-item-subtitle>
              </v-list-item-content>
          </v-list-item>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import axios from "axios";
import { mapState } from "vuex";
export default {
  name: "List",
  computed: {
    ...mapState(["username", "host","type"]),
  },
  data() {
    return {
      buyerusername:[],
      storedEvent:[],
      selectedBuyer: null,
    };
  },
  methods: {
    getEvents() {
      axios
        .put(this.host + "/list/event",{
          username:this.username,
          type:this.type
        })
        .then((response) => {
          this.storedEvent=response.data;
      });
    },
    takeBuyerUsername() {
      this.buyerusername=[]
      for(let prop in this.storedEvent){
        if (this.selectedBuyer.buyerUsername == this.storedEvent[prop].buyerUsername){
          this.buyerusername.push(this.storedEvent[prop])
        }
      }
    },
  },
  created() {
    this.getEvents()
  },
};
</script>

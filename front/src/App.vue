<template>
  <WithdrawalRequest
    :customer-id="1"
    :account-id="1"
    :portfolio-id="1"
    language="en-GB"
    currency="GBP"
  />
  <p class="mt-5 mb-3 text-muted text-center">&copy; 2021</p>
</template>

<script lang="ts">
import { Options, Vue } from 'vue-class-component';
import WithdrawalRequest from './components/WithdrawalRequest.vue';

@Options({
  components: {
    WithdrawalRequest,
  },
})
export default class App extends Vue {}
</script>

<style lang="scss">

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  margin-top: 60px;
}

.icon {
  width:58px;
  height:58px;
  padding: 5px;
}

</style>

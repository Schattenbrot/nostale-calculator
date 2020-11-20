<template>
  <div class="fairy-view-container">
    <h1>Fairy</h1>

    <b-form @submit.prevent="handleSubmit" @reset="handleReset">
      <b-form-group id="input-group-1" label="Name:" label-for="input-1">
        <b-form-input
          id="input-1"
          v-model="formData.name"
          type="text"
          required
          placeholder="Enter name"
        ></b-form-input>
      </b-form-group>

      <b-form-group id="input-group-2" label="Element:" label-for="input-2">
        <b-form-input
          id="input-2"
          type="text"
          v-model="formData.element"
          required
          placeholder="Enter element"
        ></b-form-input>
      </b-form-group>

      <b-form-group id="input-group-3" label="Max Percent:" label-for="input-3">
        <b-form-input
          id="input-3"
          type="number"
          v-model="formData.percentage"
          required
        ></b-form-input>
      </b-form-group>

      <b-button type="submit" variant="primary">Submit</b-button>
      <b-button type="reset" variant="danger">Reset</b-button>
    </b-form>
  </div>
</template>

<script>
import { mapActions } from "vuex";

export default {
  name: "Fairy",
  data() {
    return {
      formData: {
        name: "",
        element: "",
        percentage: 0,
      },
    };
  },
  methods: {
    ...mapActions(["createFairy"]),
    handleSubmit() {
      const { name, element, percentage } = this.formData;
      const payload = {
        name,
        element,
        percentage,
      };
      this.createFairy(payload);
    },
    handleReset() {
      console.log("reset pressed");
    },
  },
};
</script>

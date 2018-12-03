<template>
    <div>
        <div class="expense" v-for="expense in expenses" :key="expense.id">
            <div class="expense-data">
                <div class="expense-data-item">{{ expense.amount }}€</div>
                <div class="expense-data-item">{{ expense.description }}</div>
                <div class="expense-data-item">{{ parseDate(expense.created_at) }}</div>
            </div>
            <div class="expense-actions">
                <div class="cursor-pointer">
                    <font-awesome-icon icon="pen"></font-awesome-icon>
                </div>
                <div class="cursor-pointer ml-2" @click="deleteExpense(expense)">
                    <font-awesome-icon icon="trash"></font-awesome-icon>
                </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import { Component, Prop, Emit, Vue } from 'vue-property-decorator'
import moment from 'moment'
@Component
export default class ExpensesList extends Vue {
    @Prop() private expenses!: Array<Object>
    @Emit()
    deleteExpense (expense: Object) {
      return expense
    }
    parseDate (date: string) {
      return moment(date).format('lll')
    }
}
</script>

<style scoped>
    .expense {
        @apply bg-indigo-lightest p-5 rounded mt-2 flex justify-between;
    }
    .expense-data {
        @apply flex flex-wrap justify-around;
    }
    .expense-actions {
        @apply flex justify-around;
    }
    .expense-data-item {
        @apply ml-1;
    }
</style>

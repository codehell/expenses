<template>
    <div class="flex flex-wrap justify-center md:justify-around container mx-auto">
        <div class="expense" v-for="expense in expenses" :key="expense.id">
            <div class="expense-body">
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
            <div class="w-full flex flex-wrap text-white mt-1">
                <div class="ml-1 mt-1 bg-indigo-dark px-1 rounded-sm h-5" v-for="tag in expense.tags" :key="tag.id">{{ tag.name }}</div>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import { Component, Emit, Prop, Vue } from 'vue-property-decorator'
import moment from 'moment'

@Component
export default class ExpensesList extends Vue {
    @Prop() private expenses!: Array<Object>

    @Emit()
    deleteExpense (expense: Object) {
      return expense
    }

    public parseDate (date: string) {
      return moment(date).format('lll')
    }
}
</script>

<style scoped>

    .expense {
        @apply bg-indigo-lightest p-4 rounded mb-2 mx-1 flex flex-wrap justify-between w-full max-w-xs;
    }
    .expense-body {
        @apply flex justify-between;
    }

    .expense-data {
        @apply flex mr-2
        flex-wrap justify-around;
    }

    .expense-actions {
        @apply flex
        justify-around;
    }

    .expense-data-item {
        @apply ml-1;
    }
</style>

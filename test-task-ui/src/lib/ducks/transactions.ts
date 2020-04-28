import axios from 'axios';
import { createAction, createReducer } from '@reduxjs/toolkit';

import { AppDispatch, AppThunk } from '../store';


export interface Transaction {
    id: string;
    type: string;
    amount: number;
    effectiveDate: string;
}

const defaultTransactionsState = {
    transactions: []
};

const transactions = createAction<Transaction[]>('transactions/get')

export const getTransactions = (): AppThunk => async (
    dispatch: AppDispatch
) => {
    try {
        const request = await axios.get('http://localhost:8000/transactions');
        dispatch(transactions(request.data ? request.data : []))
    } catch (e) {
        console.log(e);
    };
};

export const transactionsReducer = createReducer(defaultTransactionsState, {
    [transactions.type]: (state, action) => {return {transactions: action.payload}}
});

export type TransactionsState = ReturnType<typeof transactionsReducer>;

export default TransactionsState;
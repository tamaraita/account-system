import React, { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import { TransactionsState, getTransactions } from '../lib/ducks/transactions';
import Table from './common/Table';


const Home = (): React.ReactElement => {
    const data = useSelector((state: TransactionsState) =>
        state.transactions
    );

    const dispatch = useDispatch();

    useEffect(() => {
    dispatch(getTransactions());
    }, [dispatch])

    return (
        <Table data={data} />
    );
};

export default Home;
use criterion::{black_box, criterion_group, criterion_main, Criterion};
use voyager_message::VoyagerMessage;
use voyager_vm::{seq, Op};

fn bench_normalize(c: &mut Criterion) {
    c.bench_function("normalize", |b| {
        b.iter_with_setup(
            || vec![mk_msg(), mk_msg(), mk_msg()],
            |op| black_box(op.into_iter().map(Op::normalize)),
        )
    });
}

fn mk_msg() -> Op<VoyagerMessage> {
    seq([
        // promise(
        //     [
        //         data(PluginMessage::new("", "")),
        //         call(FetchBlocks {
        //             chain_id: ChainId::new("chain"),
        //             start_height: Height::new_with_revision(1, 1),
        //         }),
        //         conc([
        //             noop(),
        //             data(PluginMessage::new("", "")),
        //             call(FetchBlocks {
        //                 chain_id: ChainId::new("chain"),
        //                 start_height: Height::new_with_revision(1, 1),
        //             }),
        //         ]),
        //     ],
        //     [],
        //     AggregateMsgUpdateClientsFromOrderedHeaders {
        //         chain_id: ChainId::new("chain"),
        //         counterparty_client_id: "counterparty_chain".parse().unwrap(),
        //     },
        // ),
        // seq([
        //     data(PluginMessage::new("", "")),
        //     call(FetchBlocks {
        //         chain_id: ChainId::new("chain"),
        //         start_height: Height::new_with_revision(1, 1),
        //     }),
        //     conc([
        //         noop(),
        //         data(PluginMessage::new("", "")),
        //         call(FetchBlocks {
        //             chain_id: ChainId::new("chain"),
        //             start_height: Height::new_with_revision(1, 1),
        //         }),
        //     ]),
        // ]),
    ])
}

criterion_group!(benches, bench_normalize);

criterion_main!(benches);

# API

This package adds support for `golos api`.

## State

| **ID** | **Command Name** | **Version** |
| :-: | :-: | :-: |
| **account_by_key** |
| 1 | get_key_references | **DONE** |
| **account_history** |
| 1 | get_account_history | **DONE** |
| **database_api** |
| 1 | get_account_bandwidth | **DONE** |
| 2 | get_account_count | **DONE** |
| 3 | get_accounts | **DONE** |
| 4 | get_block | **DONE** |
| 5 | get_block_header | **DONE** |
| 6 | get_chain_properties | **DONE** |
| 7 | get_config | **DONE** |
| 8 | get_conversion_requests | **DONE** |
| 9 | get_database_info | **DONE** |
| 10 | get_dynamic_global_properties | **DONE** |
| 11 | get_escrow | **DONE** |
| 12 | get_expiring_vesting_delegations | **DONE** |
| 13 | get_hardfork_version | **DONE** |
| 14 | get_next_scheduled_hardfork | **DONE** |
| 15 | get_owner_history | **DONE** |
| 16 | get_potential_signatures | **DONE** |
| 17 | get_proposed_transaction | **DONE** |
| 18 | get_recovery_request | **DONE** |
| 19 | get_required_signatures | **DONE** |
| 20 | get_savings_withdraw_from | **DONE** |
| 21 | get_savings_withdraw_to | **DONE** |
| 22 | get_transaction_hex | **DONE** |
| 23 | get_vesting_delegations | **DONE** |
| 24 | get_withdraw_routes | **DONE** |
| 25 | lookup_account_names | **DONE** |
| 26 | lookup_accounts | **DONE** |
| 27 | verify_account_authority | **DONE** |
| 28 | verify_authority | **DONE** |
| **follow** |
| 1 | get_account_reputations | **DONE** |
| 2 | get_blog | **DONE** |
| 3 | get_blog_authors | **DONE** |
| 4 | get_blog_entries | **DONE** |
| 5 | get_feed | **DONE** |
| 6 | get_feed_entries | **DONE** |
| 7 | get_follow_count | **DONE** |
| 8 | get_followers | **DONE** |
| 9 | get_following | **DONE** |
| 10 | get_reblogged_by | **DONE** |
| **market_history** |
| 1 | get_market_history | **DONE** |
| 2 | get_market_history_buckets | **DONE** |
| 3 | get_open_orders | **DONE** |
| 4 | get_order_book | **DONE** |
| 5 | get_order_book_extended | **NONE** |
| 6 | get_recent_trades | **DONE** |
| 7 | get_ticker | **DONE** |
| 8 | get_trade_history | **DONE** |
| 9 | get_volume | **DONE** |
| **network_broadcast_api** |
| 1 | broadcast_block | **NONE** |
| 2 | broadcast_transaction | **DONE** |
| 3 | broadcast_transaction_synchronous | **DONE** |
| 4 | broadcast_transaction_with_callback | **NONE** |
| **operation_history** |
| 1 | get_ops_in_block | **DONE** |
| 2 | get_transaction | **DONE** |
| **operation_history** |
| 1 | get_inbox | **DONE** |
| 2 | get_outbox | **DONE** |
| 3 | get_thread | **DONE** |
| 4 | get_settings | **DONE** |
| 5 | get_contacts_size | **RAW** |
| 6 | get_contact_info | **DONE** |
| 7 | get_contacts | **DONE** |
| **social_network** |
| 1 | get_account_votes | **DONE** |
| 2 | get_active_votes | **DONE** |
| 3 | get_all_content_replies | **DONE** |
| 4 | get_content | **DONE** |
| 5 | get_content_replies | **DONE** |
| 6 | get_replies_by_last_update | **DONE** |
| **tags** |
| 1 | get_discussions_by_active | **DONE** |
| 2 | get_discussions_by_author_before_date | **DONE** |
| 3 | get_discussions_by_blog | **DONE** |
| 4 | get_discussions_by_cashout | **DONE** |
| 5 | get_discussions_by_children | **DONE** |
| 6 | get_discussions_by_comments | **DONE** |
| 7 | get_discussions_by_created | **DONE** |
| 8 | get_discussions_by_feed | **DONE** |
| 9 | get_discussions_by_hot | **DONE** |
| 10 | get_discussions_by_payout | **DONE** |
| 11 | get_discussions_by_promoted | **DONE** |
| 12 | get_discussions_by_trending | **DONE** |
| 13 | get_discussions_by_votes | **DONE** |
| 14 | get_languages | **DONE** |
| 15 | get_tags_used_by_author | **RAW** |
| 16 | get_trending_tags | **DONE** |
| **witness_api** |
| 1 | get_active_witnesses | **DONE** |
| 2 | get_current_median_history_price | **DONE** |
| 3 | get_feed_history | **DONE** |
| 4 | get_miner_queue | **DONE** |
| 5 | get_witness_by_account | **DONE** |
| 6 | get_witness_count | **DONE** |
| 7 | get_witness_schedule | **DONE** |
| 8 | get_witnesses | **DONE** |
| 9 | get_witnesses_by_vote | **DONE** |
| 10 | lookup_witness_accounts | **DONE** |

## License

MIT, see the `LICENSE` file.

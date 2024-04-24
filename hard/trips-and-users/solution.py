import pandas as pd

def trips_and_users(trips: pd.DataFrame, users: pd.DataFrame) -> pd.DataFrame:
    # Get list of banned user IDs
    banned_user_ids = users[users['banned'] == 'Yes']['users_id']
    
    # Filter trips to include only those within the specified date range and exclude banned users
    filtered_trips = trips[
        (trips['request_at'] >= '2013-10-01') &
        (trips['request_at'] <= '2013-10-03') &
        (~trips['client_id'].isin(banned_user_ids)) &
        (~trips['driver_id'].isin(banned_user_ids))
    ]
    
    # Group by date and calculate total orders and cancellation counts
    cancellation_summary = filtered_trips.groupby('request_at').agg(
        total_orders=('status', 'count'),
        cancelled_orders=('status', lambda x: (x != 'completed').sum())
    ).reset_index()
    
    # Rename columns for clarity
    cancellation_summary.rename(columns={'request_at': 'Day'}, inplace=True)
    
    # Calculate cancellation rate and round to 2 decimal places
    cancellation_summary['Cancellation Rate'] = (
        cancellation_summary['cancelled_orders'] / cancellation_summary['total_orders']
    ).round(2)
    
    # Select and return only the relevant columns
    return cancellation_summary[['Day', 'Cancellation Rate']]
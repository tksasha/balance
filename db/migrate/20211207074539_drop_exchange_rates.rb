class DropExchangeRates < ActiveRecord::Migration[6.1]
  def change
    drop_table :exchange_rates, if_exists: true
  end
end

class AddUniqueIndexForExchangeRatesDateFromTo < ActiveRecord::Migration[6.0]
  def change
    add_index :exchange_rates, [:date, :from, :to], unique: true
  end
end

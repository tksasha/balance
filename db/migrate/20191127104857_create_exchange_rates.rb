# frozen_string_literal: true

class CreateExchangeRates < ActiveRecord::Migration[6.0]
  def change
    create_table :exchange_rates do |t|
      t.integer :from, default: 0
      t.integer :to, default: 0
      t.decimal :rate, precision: 7, scale: 5
      t.date :date

      t.timestamps
    end
  end
end

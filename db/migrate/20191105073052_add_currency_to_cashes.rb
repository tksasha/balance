# frozen_string_literal: true

class AddCurrencyToCashes < ActiveRecord::Migration[6.0]
  def change
    add_column :cashes, :currency, :integer, default: 0
  end
end

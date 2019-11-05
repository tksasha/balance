# frozen_string_literal: true

class AddCurrencyToItems < ActiveRecord::Migration[6.0]
  def change
    add_column :items, :currency, :integer, default: 0
  end
end

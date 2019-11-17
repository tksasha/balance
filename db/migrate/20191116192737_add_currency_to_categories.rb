# frozen_string_literal: true

class AddCurrencyToCategories < ActiveRecord::Migration[6.0]
  def change
    add_column :categories, :currency, :integer, default: 0
  end
end

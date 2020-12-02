class AddUniqueIndexForCashesNameAndCurrency < ActiveRecord::Migration[6.0]
  def change
    add_index :cashes, [:name, :currency], unique: true
  end
end

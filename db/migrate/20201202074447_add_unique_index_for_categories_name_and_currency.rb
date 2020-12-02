class AddUniqueIndexForCategoriesNameAndCurrency < ActiveRecord::Migration[6.0]
  def change
    add_index :categories, [:name, :currency], unique: true
  end
end

class RemoveFavoriteFromCashes < ActiveRecord::Migration[7.0]
  def change
    remove_column :cashes, :favorite, :boolean
  end
end

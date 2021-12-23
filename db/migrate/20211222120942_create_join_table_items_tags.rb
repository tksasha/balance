class CreateJoinTableItemsTags < ActiveRecord::Migration[7.0]
  def change
    create_join_table :items, :tags
  end
end

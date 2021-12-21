# frozen_string_literal: true

# == Schema Information
#
# Table name: items
#
#  id          :integer          not null, primary key
#  currency    :integer          default("uah")
#  date        :date
#  deleted_at  :time
#  description :string
#  formula     :text
#  sum         :decimal(10, 2)   not null
#  created_at  :datetime
#  updated_at  :datetime
#  category_id :integer
#
# Indexes
#
#  index_balans_items_on_date                  (date)
#  index_balans_items_on_date_and_category_id  (date,category_id)
#
# Foreign Keys
#
#  category_id  (category_id => categories.id)
#

class Consolidation < Item
  delegate :name, :income?, to: :category

  delegate :id, to: :category, prefix: true
end

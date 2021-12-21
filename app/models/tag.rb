# frozen_string_literal: true

# == Schema Information
#
# Table name: tags
#
#  id          :integer          not null, primary key
#  name        :string           not null
#  created_at  :datetime         not null
#  updated_at  :datetime         not null
#  category_id :integer          not null
#
# Indexes
#
#  index_tags_on_category_id           (category_id)
#  index_tags_on_category_id_and_name  (category_id,name) UNIQUE
#
# Foreign Keys
#
#  category_id  (category_id => categories.id)
#

class Tag < ApplicationRecord
  belongs_to :category

  validates :name, presence: true, uniqueness: { scope: :category_id }
end

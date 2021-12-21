# frozen_string_literal: true

RSpec.describe Tag, type: :model do
  it { should belong_to(:category).required }

  it { should validate_presence_of(:name) }

  describe 'validations' do
    subject { described_class.new name: Faker::Commerce.color, category: category }

    let(:category) { build :category }

    it { should validate_uniqueness_of(:name).scoped_to(:category_id) }
  end
end

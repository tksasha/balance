require 'rails_helper'

RSpec.describe CategoriesController, type: :controller do
  it { should be_an ActsAsRESTController }

  describe '#collection' do
    before { expect(Category).to receive(:order).with(:income, :name).and_return(:collection) }

    its(:collection) { should eq :collection }
  end
end

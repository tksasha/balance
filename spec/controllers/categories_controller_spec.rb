require 'rails_helper'

RSpec.describe CategoriesController, type: :controller do
  it { should be_an ActsAsRESTController }

  describe '#collection' do
    before { expect(Category).to receive(:order).with(:income).and_return(:collection) }

    its(:collection) { should eq :collection }
  end

  it_behaves_like :edit

  describe '#resource_params' do
    before { expect(subject).to receive(:params).and_return(acp category: { name: 'Drinks', income: true }) }

    its(:resource_params) { should eq permit! name: 'Drinks', income: true }
  end

  it_behaves_like :update, format: :js do
    let(:success) { -> { render_template :update } }

    let(:failure) { -> { render_template :errors } }
  end
end

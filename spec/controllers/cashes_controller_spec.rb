require 'rails_helper'

RSpec.describe CashesController, type: :controller do
  it { should be_a ActsAsRESTController }

  describe '#collection' do
    before { expect(Cash).to receive(:order).with(:name).and_return(:collection) }

    its(:collection) { should eq :collection }
  end

  describe '#resource_params' do
    before { expect(subject).to receive(:params).and_return(acp cash: { formula: '', name: '' }) }

    its(:resource_params) { should eq permit! formula: '', name: '' }
  end
end

# frozen_string_literal: true

RSpec.describe CashesController, type: :controller do
  describe '#collection' do
    context do
      before { subject.instance_variable_set :@collection, :collection }

      its(:collection) { should eq :collection }
    end

    context do
      let(:relation) { double }

      let(:params) { double }

      before { expect(subject).to receive(:params).and_return(params) }

      before { expect(Cash).to receive(:order).with(:name).and_return(relation) }

      before { expect(CashSearcher).to receive(:search).with(relation, params).and_return(:collection) }

      its(:collection) { should eq :collection }
    end
  end

  describe '#resource' do
    context do
      before { subject.instance_variable_set :@resource, :resource }

      its(:resource) { should eq :resource }
    end

    context do
      before { expect(subject).to receive(:params).and_return(id: 11) }

      before { expect(Cash).to receive(:find).with(11).and_return(:resource) }

      its(:resource) { should eq :resource }
    end
  end

  describe '#initialize_resource' do
    let(:params) { { currency: 'usd', foo: 'foo' } }

    before { expect(subject).to receive(:params).and_return(params) }

    before { expect(Cash).to receive(:new).with(currency: 'usd').and_return(:resource) }

    before { subject.send :initialize_resource }

    its(:resource) { should eq :resource }
  end

  describe '#resource_params' do
    let :params do
      acp \
        cash: {
          name: nil,
          formula: nil,
          currency: nil,
        }
    end

    before { expect(subject).to receive(:params).and_return(params) }

    its(:resource_params) { should eq params.require(:cash).permit! }
  end

  describe '#build_resource' do
    before { expect(subject).to receive(:resource_params).and_return(:resource_params) }

    before { expect(Cash).to receive(:new).with(:resource_params).and_return(:resource) }

    before { subject.send :build_resource }

    its(:resource) { should eq :resource }
  end

  describe '#set_variant' do
    context do
      before { expect(subject).to receive(:params).and_return({}) }

      after { subject.send :set_variant }

      it { expect(subject).not_to receive(:request) }
    end

    context do
      before { expect(subject).to receive(:params).and_return(report: '') }

      after { subject.send :set_variant }

      it { expect(subject).not_to receive(:request) }
    end

    context do
      before { expect(subject).to receive(:params).and_return(report: '1') }

      after { subject.send :set_variant }

      it { expect(subject).to receive_message_chain('request.variant=').with(:report) }
    end
  end
end

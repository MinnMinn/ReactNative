import React from "react";
import {
  View,
  Text,
  StyleSheet,
  FlatList,
  Image,
  TouchableOpacity,
  TextInput,
  RefreshControl,
  Alert
} from "react-native";
import LinearGradient from "react-native-linear-gradient";
import { getFoods, deleteFoodById } from '../API/Api';
import { listFoods, deleteFood } from '../actions/index';
import { connect } from 'react-redux';
import Icon from '../icon';

class All extends React.Component {

  constructor(props) {
    super(props);
    this.state = {
      data: this.props.props.foods,
      data_temp: this.props.props.foods,
      refreshing: false,
      search: ''
    }
  }

  getAll = () => {
    this.setState({
      refreshing: true
    })
    getFoods().then(foods => {
      this.setState({
        data: foods,
        data_temp: foods,
        refreshing: false
      })
    }).catch((error) => {
      this.setState({
        data: [],
        data_temp: [],
        refreshing: false
      })
    });
  }

  _rating(item) {
    let rating = [];
    for (i = 0; i < item; i++) {
      rating.push(
        <Image
          source={require("../asset/star.png")}
          style={{ width: 15, height: 15, marginRight: 3 }}
          resizeMode="cover"
        />
      )
    }
    return rating;
  }

  renderItem = ({ item }) => {
    return (
      <>
        <LinearGradient
          colors={['#009245', '#8cc631']}
          start={{ x: 0, y: 1 }} end={{ x: 1, y: 0 }}
          style={styles.item}
        >
          <View style={styles.image_container}>
            <Image
              source={{ uri: item.image }}
              style={styles.image}
            />
          </View>
          <View style={styles.content}>
            <Text style={styles.name}>{item.title}</Text>
            <View style={styles.rating}>
              {this._rating(item.avgStars)}
            </View>
            <View style={styles.price_container}>
              <View style={styles.price}>
                <Text style={styles.textPrice}>{item.price}</Text>
              </View>
            </View>
          </View>
          <TouchableOpacity
            onPress={() => this.props.props.navigation.navigate("DetailScreen", {
              item: item
            })}
            style={styles.button}>
            <Icon
              name="enter"
              color="#8cc631"
              size={20}
            />
          </TouchableOpacity>
          <TouchableOpacity
            name="delete"
            style={styles.button}
            onPress={() => this.deleteFood(item.id)}
          ><Icon
              name="cross"
              color="#8cc631"
              size={20}
            /></TouchableOpacity>
        </LinearGradient>
      </>
    )
  }

  ItemSeparatorComponent = () => {
    return (
      <View
        style={{
          height: 10
        }}
      />
    )
  }

  _search(text) {
    let data = this.state.data_temp.filter(value => {
      return value.title.toUpperCase().indexOf(text) !== -1;
    });
    this.setState({
      data: data,
      search: text
    });
  }

  async deleteFood(id) {
    await Alert.alert(
      "DELETE",
      "Delete Food",
      [
        {
          text: "Cancel",
          onPress: () => console.log("Cancel Pressed"),
          style: "cancel"
        },
        { text: "OK", onPress: () => { deleteFoodById(id), this.getAll } }
      ],
      { cancelable: false }
    );
  }

  render() {
    return (
      <View style={styles.container}>
        <View style={styles.section}>
          <TextInput
            placeholder="Search..."
            style={{ flex: 1, marginLeft: 10 }}
            value={this.state.search}
            keyboardAppearance={true}
            onChangeText={(text) => this._search(text)}
          />
          <TouchableOpacity
            onPress={() => this._search("")}
            style={{ paddingHorizontal: 10 }}>
            {/* <ion-icon name="arrow-back-circle-outline"></ion-icon> */}
          </TouchableOpacity>

        </View>
        <View style={styles.flatList}>
          <FlatList
            data={this.state.data}
            renderItem={this.renderItem}
            keyExtractor={(item, index) => index.toString()}
            ItemSeparatorComponent={this.ItemSeparatorComponent}
            showsVerticalScrollIndicator={true}
            refreshControl={
              <RefreshControl
                refreshing={this.state.refreshing}
                onRefresh={this.getAll}
              />
            }
          />
        </View>
      </View>
    )
  }
}

var styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: 'white',
    paddingBottom: 5
  },
  flatList: {
    flex: 1,
    marginTop: 10
  },
  item: {
    flex: 1,
    paddingVertical: 10,
    paddingHorizontal: 10,
    flexDirection: 'row',
    borderRadius: 10
  },
  image_container: {
    width: 90,
    height: 90
  },
  image: {
    width: '100%',
    height: '100%',
    borderWidth: 5,
    borderColor: 'white',
    borderRadius: 10
  },
  content: {
    flex: 1,
    justifyContent: 'center',
    paddingHorizontal: 10
  },
  name: {
    color: 'white',
    fontWeight: 'bold',
    fontSize: 18
  },
  rating: {
    marginTop: 5,
    flexDirection: 'row'
  },
  button: {
    width: 30,
    height: 30,
    backgroundColor: 'white',
    borderRadius: 15,
    justifyContent: 'center',
    alignItems: 'center'
  },
  price_container: {
    flexDirection: 'row',
    marginTop: 10
  },
  price: {
    backgroundColor: 'white',
    paddingVertical: 5,
    paddingHorizontal: 15,
    borderRadius: 50
  },
  textPrice: {
    color: 'green',
    fontWeight: 'bold'
  },
  section: {
    flexDirection: 'row',
    alignItems: 'center',
    paddingVertical: 5,
    paddingHorizontal: 10,
    borderRadius: 100,
    backgroundColor: '#f2f2f2',
    marginTop: 10
  }
});

function mapStateToProps(state) {
  return {

  };
}

function mapDispatchToProps(dispatch, props) {
  return {
    listFoods: (foods) => {
      dispatch(listFoods(foods));
    },
    deleteFood: (food) => {
      dispatch(deleteFood(food));
    }
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(All);